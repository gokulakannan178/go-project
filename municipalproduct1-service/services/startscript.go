package services

import (
	"fmt"
	"municipalproduct1-service/constants"
	"municipalproduct1-service/models"
)

//CreateIndexes : ""
func (s *Service) CreateIndexes(ctx *models.Context) error {
	defer ctx.Session.EndSession(ctx.CTX)
	//Property Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPROPERTY, []string{"doa", "yoa", "status", "uniqueId"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPROPERTY + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPROPERTY + " index created")
	//Propertyconfiguration Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPROPERTYCONFIGURATION, []string{"uniqueId"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPROPERTYCONFIGURATION + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPROPERTYCONFIGURATION + " index created")

	//Collection Other Charges Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONOTHERCHARGES, []string{"uniqueId"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONOTHERCHARGES + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONOTHERCHARGES + " index created")

	//Floor Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPROPERTYFLOOR, []string{"uniqueId", "propertyId", "status", "dateFrom"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPROPERTYFLOOR + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPROPERTYFLOOR + " index created")

	//Floor Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONUSAGETYPE, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONUSAGETYPE + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONUSAGETYPE + " index created")

	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONCONSTRUCTIONTYPE, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONCONSTRUCTIONTYPE + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONCONSTRUCTIONTYPE + " index created")
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONOCCUMANCYTYPE, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONOCCUMANCYTYPE + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONOCCUMANCYTYPE + " index created")
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONNONRESIDENTIALUSAGEFACTOR + " index created")
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONFLOORRATABLEAREA, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONFLOORRATABLEAREA + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONFLOORRATABLEAREA + " index created")
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONFLOORTYPE, []string{"uniqueId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONFLOORTYPE + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONFLOORTYPE + " index created")

	//AVR Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONAVR, []string{"uniqueId", "municipalityTypeId", "constructionTypeId", "roadTypeId", "usageTypeId"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONAVR + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONAVR + " index created")

	//propertypaymentfys Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPROPERTYPAYMENTFY, []string{"uniqueId", "propertyId", "status"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPROPERTYPAYMENTFY + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPROPERTYPAYMENTFY + " index created")

	//FYs Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONFINANCIALYEAR, []string{"uniqueId", "order", "to", "from"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONFINANCIALYEAR + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONFINANCIALYEAR + " index created")

	//Vacant Land Rate Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONVACANTLANDRATE, []string{"uniqueId", "doe", "roadTypeId", "municipalityTypeId"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONVACANTLANDRATE + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONVACANTLANDRATE + " index created")

	//Property tax Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPROPERTYTAX, []string{"uniqueId", "doe", "to", "from"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPROPERTYTAX + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPROPERTYTAX + " index created")

	//Property tax Collection
	if err := s.Daos.EnsureIndex(ctx, constants.COLLECTIONPENALTY, []string{"uniqueId", "doe", "to", "from"}); err != nil {
		fmt.Println("err in creating " + constants.COLLECTIONPENALTY + "index" + err.Error())
	}
	fmt.Println(constants.COLLECTIONPENALTY + " index created")

	return nil
}
