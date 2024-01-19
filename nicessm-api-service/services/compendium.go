package services

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"log"
	"nicessm-api-service/constants"
	"nicessm-api-service/models"
	"os"
	"strings"
	"time"

	"github.com/tenkoh/go-docc"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) SaveCompendium(ctx *models.Context, content *models.Compendium) error {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	content.Status = constants.MARKETSTATUSACTIVE
	content.ActiveStatus = true
	//market.District = primitive.ObjectID
	t := time.Now()
	//created := models.Created{}
	//created.On = &t
	//created.By = constants.SYSTEM
	log.Println("b4 organisation.created")
	content.DateCreated = &t
	//content.DateReviewed = &t
	log.Println("b4 organisation.created")
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		dberr := s.Daos.SaveCompendium(ctx, content)
		if dberr != nil {
			if err1 := ctx.Session.AbortTransaction(sc); err1 != nil {
				log.Println("err in abort")
				return errors.New("Transaction Aborted with error" + err1.Error())
			}
			log.Println("err in abort out")
			return errors.New("Transaction Aborted - " + dberr.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateCompendium(ctx *models.Context, content *models.Compendium) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.UpdateCompendium(ctx, content)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) EnableCompendium(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.EnableCompendium(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) DisableCompendium(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DisableCompendium(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteCompendium(ctx *models.Context, UniqueID string) error {
	if err := ctx.Session.StartTransaction(); err != nil {
		return err
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {

		err := s.Daos.DeleteCompendium(ctx, UniqueID)
		if err != nil {
			if err = ctx.Session.AbortTransaction(sc); err != nil {
				return errors.New("Transaction Aborted with error" + err.Error())
			}
			return errors.New("Transaction Aborted - " + err.Error())
		}
		return nil

	}); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetSingleCompendium(ctx *models.Context, UniqueID string) (*models.RefCompendium, error) {
	content, err := s.Daos.GetSingleCompendium(ctx, UniqueID)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (s *Service) FilterCompendium(ctx *models.Context, compendiumfilter *models.CompendiumFilter, pagination *models.Pagination) (content []models.RefCompendium, err error) {
	return s.Daos.FilterCompendium(ctx, compendiumfilter, pagination)
}
func (s *Service) CompendiumUploadWord(ctx *models.Context, file *models.CompendiumFileUpload) ([]string, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var ResultStrings []string
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("files===>", file.File)
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.COMPENDIUMFILE)
		file := fmt.Sprintf("%v%v", documentUrl, file.File)
		fmt.Println("filesUri===>", file)
		r, err := docc.NewReader(file)
		if err != nil {
			panic(err)
		}
		defer r.Close()
		ps, _ := r.ReadAll()
		for _, p := range ps {
			//strings.Split(p, "/n")
			if p != "" {
				ResultStrings = append(ResultStrings, p)

			}
		}

		return nil

	}); err != nil {
		return nil, err
	}
	return ResultStrings, nil
}

func (s *Service) CompendiumUploadWordV2(ctx *models.Context, f *models.CompendiumFileUpload) ([]models.CompendiumData, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var ResultStrings []string
	var compediumData []models.CompendiumData
	t := time.Now()
	var uniqueid string
	if int(t.Month()) < 10 {
		if t.Day() < 10 {
			uniqueid = fmt.Sprintf("%v0%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v0%v%v", t.Year(), int(t.Month()), t.Day())
		}
	} else {
		if t.Day() < 0 {
			uniqueid = fmt.Sprintf("%v%v0%v", t.Year(), int(t.Month()), t.Day())
		} else {
			uniqueid = fmt.Sprintf("%v%v%v", t.Year(), int(t.Month()), t.Day())
		}
	}
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		pc, err := s.Daos.GetactiveProductConfig(ctx, true)
		if err != nil {
			return err
		}
		fmt.Println("files===>", f.File)
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.COMPENDIUMFILE)
		file := fmt.Sprintf("%v%v", documentUrl, f.File)
		fmt.Println("filesUri===>", file)
		r, err := docc.NewReader(file)
		if err != nil {
			panic(err)
		}
		defer r.Close()
		var compedium = models.CompendiumData{}
		var tag bool = false
		fmt.Println(tag)
		for {
			p, err := r.Read()
			if err == io.EOF {
				if len(compediumData) > 0 {
					if len(compediumData[0].ContentStrings) > 0 {
						compediumData[0].Title = compediumData[0].ContentStrings[0]
						compediumData[0].ContentStrings = compediumData[0].ContentStrings[1:]
					}
				}
				if f.DoSave {
					compedium := new(models.Compendium)
					compedium.File = file
					compedium.Status = constants.COMPENDIUMSTATUSACTIVE
					compedium.DateCreated = &t
					err := s.Daos.SaveCompendium(ctx, compedium)
					if err != nil {
						return err
					}
					for _, v := range compediumData {
						content := ""
						for _, v2 := range v.ContentStrings {
							content = content + v2
						}
						f.Content.Content = content
						f.Content.ContentTitle = v.Title
						f.Content.Status = "A"
						f.Content.Organisation = pc.Orgnisation.OrgnisationID
						f.Content.Project = pc.Project.ProjectID
						f.Content.ID = primitive.NewObjectID()
						f.Content.Type = "S"

						f.Content.RecordId = fmt.Sprintf("%v%v_%v", "S", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
						if f.Content.DateCreated == nil {
							f.Content.DateCreated = &t
						}
						f.Content.DateReviewed = &t
						f.Content.Tag = v.Tag
						f.Content.Tentativedate = v.Date
						f.Content.Compendium = compedium.ID
						if err := s.Daos.SaveContent(ctx, &f.Content); err != nil {
							return err
						}
						f.Content.Type = "P"
						f.Content.Content = fmt.Sprintf("<h1 style='text-align:center;color:red'>%v</h1><p>%v</p><h6>%v</h6><h6>%v</h6>", f.Content.ContentTitle, f.Content.Content, f.Content.Tag, f.Content.Tentativedate)
						f.Content.RecordId = fmt.Sprintf("%v%v_%v", "P", uniqueid, s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
						f.Content.ID = primitive.NewObjectID()
						if err := s.Daos.SaveContent(ctx, &f.Content); err != nil {
							return err
						}
						// f.Content.Type = "D"
						// f.Content.RecordId = fmt.Sprintf("%v%v_%v", uniqueid, "D", s.Daos.GetUniqueID(ctx, constants.COLLECTIONCONTENT))
						// f.Content.ID = primitive.NewObjectID()
						// uri, err := s.CompendiumPdfDocument(ctx, f.Content.Content, f.Content.RecordId)
						// if err != nil {
						// 	return err
						// }
						// f.Content.Content = uri
						// fmt.Println("document==>", f.Content.Content)
						// if err := s.Daos.SaveContent(ctx, &f.Content); err != nil {
						// 	return err
						// }
					}

				}
				return nil
			} else if err != nil {
				panic(err)
			}
			if isTag := strings.Contains(p, "Tags:"); isTag {
				compedium.Tag = p

				tag = true
			} else if isDate := strings.Contains(p, "Tentative date:"); isDate {
				compedium.Date = p
				// compediumData = append(compediumData, compedium)
				// compedium = models.CompendiumData{}
				// tag = false
				continue

			} else {
				if p == "" {
					continue
				}
				if tag {
					compediumData = append(compediumData, compedium)
					compedium = models.CompendiumData{}
					// compedium.ContentStrings = []string{}
					compedium.Title = p
					tag = false
					continue
				}
				compedium.ContentStrings = append(compedium.ContentStrings, p)
				ResultStrings = append(ResultStrings, p)
			}

			// do something with p:string
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return compediumData, nil
}

func (s *Service) CompendiumUploadWordV3(ctx *models.Context, file *models.CompendiumFileUpload) ([]string, error) {
	log.Println("transaction start")
	//Start Transaction
	if err := ctx.Session.StartTransaction(); err != nil {
		return nil, err
	}
	var ResultStrings []string
	defer ctx.Session.EndSession(ctx.CTX)
	if err := mongo.WithSession(ctx.CTX, ctx.Session, func(sc mongo.SessionContext) error {
		fmt.Println("files===>", file.File)
		documentUrl := s.Shared.Config.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.COMPENDIUMFILE)
		file := fmt.Sprintf("%v%v", documentUrl, file.File)
		fmt.Println("filesUri===>", file)
		r, err := docc.NewReader(file)
		if err != nil {
			panic(err)
		}
		defer r.Close()
		for {
			p, err := r.Read()
			if err == io.EOF {
				return nil
			} else if err != nil {
				panic(err)
			}
			// do something with p:string
			ResultStrings = append(ResultStrings, p)

		}

		return nil

	}); err != nil {
		return nil, err
	}
	return ResultStrings, nil
}
func (s *Service) CompendiumPdfDocument(ctx *models.Context, content string, recordId string) (string, error) {
	r := NewRequestPdf("")
	byts, err := r.ContentDocumePdfGeneration(content)
	if err != nil {
		return "", err
	}
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	st := fmt.Sprintf("%X", b)
	docStart := s.ConfigReader.GetString(s.Shared.GetCmdArg(constants.ENV) + "." + constants.DOCLOC)
	scenario := "kmdocument"
	fileuri := docStart + scenario + "/" + st + recordId + ".pdf"
	responseURI := constants.DEFAULTFILEURL + scenario + "/" + st + recordId + ".pdf"
	fmt.Println("fileuri=", fileuri)
	// open output file
	// file := fmt.Sprintf("%v", fileuri)
	fo, err := os.Create(fileuri)
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	if _, err := fo.Write(byts); err != nil {
		panic(err)
	}
	return responseURI, nil
}
