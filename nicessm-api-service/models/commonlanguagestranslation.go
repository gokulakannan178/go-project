package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommonLanguageTranslationss struct {
	ID           primitive.ObjectID `json:"id" form:"id," bson:"_id,omitempty"`
	LanguageType string             `json:"languageType" form:"languageType" bson:"languageType,omitempty"`
	Dashboard    struct {
		Dashboard      string `json:"dashboard" bson:"dashboard,omitempty"`
		AdminDashboard string `json:"adminDashboard" bson:"adminDashboard,omitempty"`
		Home           string `json:"home" bson:"home,omitempty"`
		Content        struct {
			Sms          string `json:"sms" bson:"sms,omitempty"`
			Voice        string `json:"voice" bson:"voice,omitempty"`
			Video        string `json:"video" bson:"video,omitempty"`
			Poster       string `json:"poster" bson:"poster,omitempty"`
			Document     string `json:"document" bson:"document,omitempty"`
			UnReviewed   string `json:"unReviewed" bson:"unReviewed,omitempty"`
			Reviewed     string `json:"reviewed" bson:"reviewed,omitempty"`
			Rejected     string `json:"rejected" bson:"rejected,omitempty"`
			Deleted      string `json:"deleted" bson:"deleted,omitempty"`
			First        string `json:"first" bson:"first,omitempty"`
			Last         string `json:"last" bson:"last,omitempty"`
			Next         string `json:"next" bson:"next,omitempty"`
			Previous     string `json:"previous" bson:"previous,omitempty"`
			TotalRecords string `json:"totalRecords" bson:"totalRecords,omitempty"`
			Content      string `json:"content" bson:"content,omitempty"`
		} `json:"content" bson:"content,omitempty"`
		Users struct {
			CallCenterAgent     string `json:"callCenterAgent" bson:"callCenterAgent,omitempty"`
			ContentCreator      string `json:"contentCreator" bson:"contentCreator,omitempty"`
			ContentManager      string `json:"contentManager" bson:"contentManager,omitempty"`
			ContentProvider     string `json:"contentProvider" bson:"contentProvider,omitempty"`
			ContentDisseminator string `json:"contentDisseminator" bson:"contentDisseminator,omitempty"`
			FieldAgent          string `json:"fieldAgent" bson:"fieldAgent,omitempty"`
			LanguageTranslator  string `json:"languageTranslator" bson:"languageTranslator,omitempty"`
			TranslationApprover string `json:"translationApprover" bson:"translationApprover,omitempty"`
			Management          string `json:"management" bson:"management,omitempty"`
			DistrictAdmin       string `json:"districtAdmin" bson:"districtAdmin,omitempty"`
			SubjectMatteExpert  string `json:"subjectMatteExpert" bson:"subjectMatteExpert,omitempty"`
			SystemAdmin         string `json:"systemAdmin" bson:"systemAdmin,omitempty"`
			VisitorViewer       string `json:"visitorViewer" bson:"visitorViewer,omitempty"`
			AllUsers            string `json:"allUsers" bson:"allUsers,omitempty"`
			Trainer             string `json:"trainer" bson:"trainer,omitempty"`
			GuestUsers          string `json:"guestUsers" bson:"guestUsers,omitempty"`
			FieldAgentLead      string `json:"fieldAgentLead" bson:"fieldAgentLead,omitempty"`
			User                string `json:"users" bson:"users,omitempty"`
		} `json:"users" bson:"users,omitempty"`
		Farmer struct {
			ActiveFarmers   string `json:"activeFarmers" bson:"activeFarmers,omitempty"`
			InactiveFarmers string `json:"inactiveFarmers" bson:"inactiveFarmers,omitempty"`
			Farmer          string `json:"farmer" bson:"farmer,omitempty"`
		} `json:"farmer" bson:"farmer,omitempty"`
		Query struct {
			UnresolvedQueries string `json:"unresolvedQueries" bson:"unresolvedQueries,omitempty"`
			AssignedQueries   string `json:"assignedQueries" bson:"assignedQueries,omitempty"`
			ResolvedQueries   string `json:"resolvedQueries" bson:"resolvedQueries,omitempty"`
			Query             string `json:"query" bson:"query,omitempty"`
		} `json:"query" bson:"query,omitempty"`
		Weather struct {
			WeatherData        string `json:"weatherData" bson:"weatherData,omitempty"`
			UPDATEASOF         string `json:"updatesOf" bson:"updatesOf,omitempty"`
			StartDate          string `json:"startDate" bson:"startDate,omitempty"`
			LastDate           string `json:"lastDate" bson:"lastDate,omitempty"`
			Search             string `json:"search" bson:"search,omitempty"`
			Reset              string `json:"reset" bson:"reset,omitempty"`
			Rain               string `json:"rain" bson:"rain,omitempty"`
			MinimumTemperature string `json:"minimumTemperature" bson:"minimumTemperature,omitempty"`
			MaximumTemperature string `json:"maximumTemperature" bson:"maximumTemperature,omitempty"`
			Damp               string `json:"damp" bson:"damp,omitempty"`
			WindSpeed          string `json:"windSpeed" bson:"windSpeed,omitempty"`
			First              string `json:"first" bson:"first,omitempty"`
			Last               string `json:"last" bson:"last,omitempty"`
			Next               string `json:"next" bson:"next,omitempty"`
			Previous           string `json:"previous" bson:"previous,omitempty"`
			State              string `json:"state" bson:"state,omitempty"`
		} `json:"weather" bson:"weather,omitempty"`
		Profile struct {
			Profile        string `json:"profile" bson:"profile,omitempty"`
			Access         string `json:"access" bson:"access,omitempty"`
			Org            string `json:"org" bson:"org,omitempty"`
			ChangePassword string `json:"changePassword" bson:"changePassword,omitempty"`
			Logout         string `json:"logout" bson:"logout,omitempty"`
		} `json:"profile" bson:"profile,omitempty"`
	} `json:"dashboard" bson:"dashboard,omitempty"`
	Master struct {
		Master   string `json:"master" bson:"master,omitempty"`
		Language struct {
			Language    string `json:"language" bson:"language,omitempty"`
			Name        string `json:"name" bson:"name,omitempty"`
			Status      string `json:"status" bson:"status,omitempty"`
			Controls    string `json:"controls" bson:"controls,omitempty"`
			Close       string `json:"close" bson:"close,omitempty"`
			Save        string `json:"save" bson:"save,omitempty"`
			NewLanguage string `json:"newLanguage" bson:"newLanguage,omitempty"`
		} `json:"language" bson:"language,omitempty"`
		SoilType struct {
			SoilType    string `json:"soilType" bson:"soilType,omitempty"`
			Name        string `json:"name" bson:"name,omitempty"`
			Status      string `json:"status" bson:"status,omitempty"`
			Controls    string `json:"controls" bson:"controls,omitempty"`
			Close       string `json:"close" bson:"close,omitempty"`
			Save        string `json:"save" bson:"save,omitempty"`
			NewSoilType string `json:"newSoilType" bson:"newSoilType,omitempty"`
		} `json:"soilType" bson:"soilType,omitempty"`
		Location struct {
			Location struct {
				Location  string `json:"location" bson:"location,omitempty"`
				Name      string `json:"name" bson:"name,omitempty"`
				Status    string `json:"status" bson:"status,omitempty"`
				Controls  string `json:"controls" bson:"controls,omitempty"`
				Close     string `json:"close" bson:"close,omitempty"`
				Save      string `json:"save" bson:"save,omitempty"`
				NewState  string `json:"newState" bson:"newState,omitempty"`
				Languages string `json:"languages" bson:"languages,omitempty"`
			} `json:"location" bson:"location,omitempty"`
			Imd struct {
				Imd      string `json:"imd" bson:"imd,omitempty"`
				URL      string `json:"url" bson:"url,omitempty"`
				Level    string `json:"level" bson:"level,omitempty"`
				Controls string `json:"controls" bson:"controls,omitempty"`
			} `json:"imd" bson:"imd,omitempty"`
			Cluster struct {
				Cluster       string `json:"cluster" bson:"cluster,omitempty"`
				Name          string `json:"name" bson:"name,omitempty"`
				Status        string `json:"status" bson:"status,omitempty"`
				Controls      string `json:"controls" bson:"controls,omitempty"`
				Close         string `json:"close" bson:"close,omitempty"`
				Save          string `json:"save" bson:"save,omitempty"`
				NewCluster    string `json:"newCluster" bson:"newCluster,omitempty"`
				State         string `json:"state" bson:"state,omitempty"`
				District      string `json:"district" bson:"district,omitempty"`
				Village       string `json:"village" bson:"village,omitempty"`
				Block         string `json:"block" bson:"block,omitempty"`
				GramPanchayat string `json:"gramPanchayat" bson:"gramPanchayat,omitempty"`
			} `json:"cluster" bson:"cluster,omitempty"`
			CommonLand struct {
				CommonLand    string `json:"commonLand" bson:"commonLand,omitempty"`
				ParcelNumber  string `json:"parcelNumber" bson:"parcelNumber,omitempty"`
				Status        string `json:"status" bson:"status,omitempty"`
				Controls      string `json:"controls" bson:"controls,omitempty"`
				Close         string `json:"close" bson:"close,omitempty"`
				Save          string `json:"save" bson:"save,omitempty"`
				NewCommonLand string `json:"newCommonLand" bson:"newCommonLand,omitempty"`
				State         string `json:"state" bson:"state,omitempty"`
				District      string `json:"district" bson:"district,omitempty"`
				Village       string `json:"village" bson:"village,omitempty"`
				Block         string `json:"block" bson:"block,omitempty"`
				GramPanchayat string `json:"gramPanchayat" bson:"gramPanchayat,omitempty"`
				KhasraNo      string `json:"khasraNo" bson:"khasraNo,omitempty"`
				Type          string `json:"type" bson:"type,omitempty"`
				OwnerShip     string `json:"ownerShip" bson:"ownerShip,omitempty"`
				Latitude      string `json:"latitude" bson:"latitude,omitempty"`
				Longitude     string `json:"longitude" bson:"longitude ,omitempty"`
			} `json:"commonLand" bson:"commonLand,omitempty"`
		} `json:"location" bson:"location,omitempty"`
		Market struct {
			ViewMarket   string `json:"viewMarket" bson:"viewMarket,omitempty"`
			EditMarket   string `json:"editMarket" bson:"editMarket,omitempty"`
			Market       string `json:"market" bson:"market,omitempty"`
			Status       string `json:"status" bson:"status,omitempty"`
			Controls     string `json:"controls" bson:"controls,omitempty"`
			Close        string `json:"close" bson:"close,omitempty"`
			AddMarket    string `json:"addMarket" bson:"addMarket,omitempty"`
			CreateMarket string `json:"createMarket" bson:"createMarket,omitempty"`
			Name         string `json:"name" bson:"name,omitempty"`
			Desc         string `json:"description" bson:"description,omitempty"`
			Level        string `json:"level" bson:"level,omitempty"`
			AddressLine1 string `json:"addressLine1" bson:"addressLine1,omitempty"`
			AddressLine2 string `json:"addressLine2" bson:"addressLine2,omitempty"`
			Latitude     string `json:"latitude" bson:"latitude,omitempty"`
			Longitude    string `json:"longitude" bson:"longitude ,omitempty"`
		} `json:"market" bson:"market,omitempty"`
		Asset struct {
			Asset    string `json:"asset" bson:"asset,omitempty"`
			Name     string `json:"name" bson:"name,omitempty"`
			Status   string `json:"status" bson:"status,omitempty"`
			Controls string `json:"controls" bson:"controls,omitempty"`
			Close    string `json:"close" bson:"close,omitempty"`
			Save     string `json:"save" bson:"save,omitempty"`
			NewAsset string `json:"newAsset" bson:"newAsset,omitempty"`
		} `json:"asset" bson:"asset,omitempty"`
		KnowledgeDomain struct {
			KnowledgeDomain    string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
			Name               string `json:"name" bson:"name,omitempty"`
			Status             string `json:"status" bson:"status,omitempty"`
			Controls           string `json:"controls" bson:"controls,omitempty"`
			Close              string `json:"close" bson:"close,omitempty"`
			Save               string `json:"save" bson:"save,omitempty"`
			Desc               string `json:"description" bson:"description,omitempty"`
			NewKnowledgeDomain string `json:"newKnowledgeDomain" bson:"newKnowledgeDomain,omitempty"`
			SubDomain          string `json:"subDomain" bson:"subDomain,omitempty"`
			Topics             struct {
				Topics             string `json:"topics" bson:"topics,omitempty"`
				TimeApplicableType string `json:"timeApplicableType" bson:"timeApplicableType,omitempty"`
				Independent        string `json:"independent" bson:"independent,omitempty"`
				Monthofeveryyear   string `json:"monthofeveryyear" bson:"monthofeveryyear,omitempty"`
				DateRange          string `json:"dateRange" bson:"dateRange,omitempty"`
				SpecificDate       string `json:"specificDate" bson:"specificDate,omitempty"`
				Classfication      string `json:"classfication" bson:"classfication,omitempty"`
				Cause              string `json:"cause" bson:"cause,omitempty"`
				SoilType           string `json:"soilType" bson:"soilType,omitempty"`
				Commodity          string `json:"commodity" bson:"commodity,omitempty"`
				CommodityStage     string `json:"commodityStage" bson:"commodityStage,omitempty"`
				SubVariety         string `json:"subVariety" bson:"subVariety,omitempty"`
				Causative          string `json:"causative" bson:"causative,omitempty"`
				Irriagtion         string `json:"irriagtion" bson:"irriagtion,omitempty"`
				CauseType          string `json:"causeType" bson:"causeType,omitempty"`
				Season             string `json:"season" bson:"season,omitempty"`
				Market             string `json:"market" bson:"market,omitempty"`
				Category           string `json:"category" bson:"category,omitempty"`
				Variety            string `json:"variety" bson:"variety,omitempty"`
				SubCategory        string `json:"subCategory" bson:"subCategory,omitempty"`
				SoilData           string `json:"soilData" bson:"soilData,omitempty"`
			} `json:"topics" bson:"topics,omitempty"`
			SubTopic string `json:"subTopic" bson:"subTopic,omitempty"`
		} `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
		CropCategory struct {
			CropCategory    string `json:"cropCategory" bson:"cropCategory,omitempty"`
			Name            string `json:"name" bson:"name,omitempty"`
			Status          string `json:"status" bson:"status,omitempty"`
			Controls        string `json:"controls" bson:"controls,omitempty"`
			Close           string `json:"close" bson:"close,omitempty"`
			Save            string `json:"save" bson:"save,omitempty"`
			Classification  string `json:"classification" bson:"classification,omitempty"`
			NewCropCategory string `json:"newCropCategory" bson:"newCropCategory,omitempty"`
		} `json:"cropCategory" bson:"cropCategory,omitempty"`
		LiveStockCategory struct {
			LiveStockCategory string `json:"liveStockCategory" bson:"liveStockCategory,omitempty"`
			SubCategory       struct {
				Name            string `json:"name" bson:"name,omitempty"`
				Status          string `json:"status" bson:"status,omitempty"`
				Controls        string `json:"controls" bson:"controls,omitempty"`
				Close           string `json:"close" bson:"close,omitempty"`
				Save            string `json:"save" bson:"save,omitempty"`
				NewCropCategory string `json:"newCropCategory" bson:"newCropCategory,omitempty"`
			} `json:"subCategory" bson:"subCategory,omitempty"`
			Name            string `json:"name" bson:"name,omitempty"`
			Status          string `json:"status" bson:"status,omitempty"`
			Controls        string `json:"controls" bson:"controls,omitempty"`
			Close           string `json:"close" bson:"close,omitempty"`
			Save            string `json:"save" bson:"save,omitempty"`
			Classification  string `json:"classification" bson:"classification,omitempty"`
			NewCropCategory string `json:"newCropCategory" bson:"newCropCategory,omitempty"`
		} `json:"liveStockCategory" bson:"liveStockCategory,omitempty"`
		AidCategory struct {
			AidCategory    string `json:"aidCategory" bson:"aidCategory,omitempty"`
			Name           string `json:"name" bson:"name,omitempty"`
			Status         string `json:"status" bson:"status,omitempty"`
			Controls       string `json:"controls" bson:"controls,omitempty"`
			Close          string `json:"close" bson:"close,omitempty"`
			Save           string `json:"save" bson:"save,omitempty"`
			NewAidCategory string `json:"newAidCategory" bson:"newAidCategory,omitempty"`
		} `json:"aidCategory" bson:"aidCategory,omitempty"`
		AidLocation struct {
			AidLocation    string `json:"aidLocation" bson:"aidLocation,omitempty"`
			Name           string `json:"name" bson:"name,omitempty"`
			Category       string `json:"category" bson:"category,omitempty"`
			Status         string `json:"status" bson:"status,omitempty"`
			Controls       string `json:"controls" bson:"controls,omitempty"`
			Close          string `json:"close" bson:"close,omitempty"`
			Save           string `json:"save" bson:"save,omitempty"`
			NewAidLocation string `json:"newAidLocation" bson:"newAidLocation,omitempty"`
			Desc           string `json:"description" bson:"description,omitempty"`
			Phone          string `json:"phone" bson:"phone,omitempty"`
			Address        string `json:"address" bson:"address,omitempty"`
			Latitude       string `json:"latitude" bson:"latitude,omitempty"`
			Longitude      string `json:"longitude" bson:"longitude ,omitempty"`
		} `json:"aidLocation" bson:"aidLocation,omitempty"`
		Season struct {
			Season    string `json:"season" bson:"season,omitempty"`
			Name      string `json:"name" bson:"name,omitempty"`
			Status    string `json:"status" bson:"status,omitempty"`
			Controls  string `json:"controls" bson:"controls,omitempty"`
			Close     string `json:"close" bson:"close,omitempty"`
			Save      string `json:"save" bson:"save,omitempty"`
			StartDate string `json:"startDate" bson:"startDate,omitempty"`
			EndDate   string `json:"endDate" bson:"endDate,omitempty"`
			State     string `json:"state" bson:"state,omitempty"`
			NewSeason string `json:"newSeason" bson:"newSeason,omitempty"`
		} `json:"season" bson:"season,omitempty"`
		Crop struct {
			Crop         string `json:"crop" bson:"crop,omitempty"`
			CommonName   string `json:"commonName" bson:"commonName,omitempty"`
			Status       string `json:"status" bson:"status,omitempty"`
			Controls     string `json:"controls" bson:"controls,omitempty"`
			Close        string `json:"close" bson:"close,omitempty"`
			Save         string `json:"save" bson:"save,omitempty"`
			SpecificName string `json:"specificName" bson:"specificName,omitempty"`
			NewCrop      string `json:"newCrop" bson:"newCrop,omitempty"`
			Category     string `json:"category" bson:"category,omitempty"`
			SubCategory  string `json:"subCategory" bson:"subCategory,omitempty"`
			Diseases     struct {
				Name        string `json:"name" bson:"name,omitempty"`
				Diseases    string `json:"diseases" bson:"diseases,omitempty"`
				Controls    string `json:"controls" bson:"controls,omitempty"`
				Close       string `json:"close" bson:"close,omitempty"`
				Save        string `json:"save" bson:"save,omitempty"`
				NewDiseases string `json:"newDiseases" bson:"newDiseases,omitempty"`
			} `json:"diseases" bson:"diseases,omitempty"`
			Insects struct {
				Name       string `json:"name" bson:"name,omitempty"`
				NewInsects string `json:"newInsects" bson:"newInsects,omitempty"`
				Insects    string `json:"insects" bson:"insects,omitempty"`
				Controls   string `json:"controls" bson:"controls,omitempty"`
				Close      string `json:"close" bson:"close,omitempty"`
				Save       string `json:"save" bson:"save,omitempty"`
			} `json:"insects" bson:"insects,omitempty"`
			Varities struct {
				Name        string `json:"name" bson:"name,omitempty"`
				Varities    string `json:"varities" bson:"varities,omitempty"`
				Status      string `json:"status" bson:"status,omitempty"`
				Controls    string `json:"controls" bson:"controls,omitempty"`
				Close       string `json:"close" bson:"close,omitempty"`
				Save        string `json:"save" bson:"save,omitempty"`
				NewVarities string `json:"newVarities" bson:"newVarities,omitempty"`
				SubVarities string `json:"subVarities" bson:"subVarities,omitempty"`
			} `json:"varities" bson:"varities,omitempty"`
			Stages struct {
				Name          string `json:"name" bson:"name,omitempty"`
				Stages        string `json:"stages" bson:"stages,omitempty"`
				NewStages     string `json:"newStages" bson:"newStages,omitempty"`
				Status        string `json:"status" bson:"status,omitempty"`
				Controls      string `json:"controls" bson:"controls,omitempty"`
				Close         string `json:"close" bson:"close,omitempty"`
				Save          string `json:"save" bson:"save,omitempty"`
				SquenceNumber string `json:"squenceNumber" bson:"squenceNumber,omitempty"`
			} `json:"stages" bson:"stages,omitempty"`
		} `json:"crop" bson:"crop,omitempty"`
		LiveStock struct {
			LiveStock    string `json:"liveStock" bson:"liveStock,omitempty"`
			CommonName   string `json:"commonName" bson:"commonName,omitempty"`
			Status       string `json:"status" bson:"status,omitempty"`
			Controls     string `json:"controls" bson:"controls,omitempty"`
			Close        string `json:"close" bson:"close,omitempty"`
			Save         string `json:"save" bson:"save,omitempty"`
			SpecificName string `json:"specificName" bson:"specificName,omitempty"`
			NewLiveStock string `json:"newLiveStock" bson:"newLiveStock,omitempty"`
			Category     string `json:"category" bson:"category,omitempty"`
			SubCategory  string `json:"subCategory" bson:"subCategory,omitempty"`
			Diseases     struct {
				Name        string `json:"name" bson:"name,omitempty"`
				Diseases    string `json:"diseases" bson:"diseases,omitempty"`
				Controls    string `json:"controls" bson:"controls,omitempty"`
				Close       string `json:"close" bson:"close,omitempty"`
				Save        string `json:"save" bson:"save,omitempty"`
				NewDiseases string `json:"newDiseases" bson:"newDiseases,omitempty"`
			} `json:"diseases" bson:"diseases,omitempty"`
			Varities struct {
				Name        string `json:"name" bson:"name,omitempty"`
				Varities    string `json:"varities" bson:"varities,omitempty"`
				Status      string `json:"status" bson:"status,omitempty"`
				Controls    string `json:"controls" bson:"controls,omitempty"`
				Close       string `json:"close" bson:"close,omitempty"`
				Save        string `json:"save" bson:"save,omitempty"`
				NewVarities string `json:"newVarities" bson:"newVarities,omitempty"`
			} `json:"varities" bson:"varities,omitempty"`
			Stages struct {
				Name          string `json:"name" bson:"name,omitempty"`
				Stages        string `json:"stages" bson:"stages,omitempty"`
				NewStages     string `json:"newStages" bson:"newStages,omitempty"`
				Status        string `json:"status" bson:"status,omitempty"`
				Controls      string `json:"controls" bson:"controls,omitempty"`
				Close         string `json:"close" bson:"close,omitempty"`
				Save          string `json:"save" bson:"save,omitempty"`
				SquenceNumber string `json:"squenceNumber" bson:"squenceNumber,omitempty"`
			} `json:"stages" bson:"stages,omitempty"`
		} `json:"liveStock" bson:"liveStock,omitempty"`
		Insect struct {
			Insect    string `json:"insect" bson:"insect,omitempty"`
			Name      string `json:"name" bson:"name,omitempty"`
			Status    string `json:"status" bson:"status,omitempty"`
			Controls  string `json:"controls" bson:"controls,omitempty"`
			Close     string `json:"close" bson:"close,omitempty"`
			Save      string `json:"save" bson:"save,omitempty"`
			NewInsect string `json:"newInsect" bson:"newInsect,omitempty"`
		} `json:"insect" bson:"insect,omitempty"`
		BlockSchedule struct {
			BlockSchedule string `json:"blockSchedule" bson:"blockSchedule,omitempty"`
			Season        string `json:"season" bson:"season,omitempty"`
			Crop          string `json:"crop" bson:"crop,omitempty"`
			Variety       string `json:"variety" bson:"variety,omitempty"`
			Controls      string `json:"controls" bson:"controls,omitempty"`
			TotalDays     string `json:"totalDays" bson:"totalDays,omitempty"`
			Filter        struct {
				State    string `json:"state" bson:"state,omitempty"`
				District string `json:"district" bson:"district,omitempty"`
				Block    string `json:"block" bson:"block,omitempty"`
				Crop     string `json:"crop" bson:"crop,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"blockSchedule" bson:"blockSchedule,omitempty"`
		BlockCrop struct {
			BlockCrop           string `json:"blockCrop" bson:"blockCrop,omitempty"`
			Commodity           string `json:"commodity" bson:"commodity,omitempty"`
			NameInLocalLanguage string `json:"nameInLocalLanguage" bson:"nameInLocalLanguage,omitempty"`
			BlockName           string `json:"blockName" bson:"blockName,omitempty"`
			Controls            string `json:"controls" bson:"controls,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Close               string `json:"close" bson:"close,omitempty"`
			AddBlockCrop        string `json:"addBlockCrop" bson:"addBlockCrop,omitempty"`
			CreateBlockCrop     string `json:"createBlockCrop" bson:"createBlockCrop,omitempty"`
			Filter              struct {
				State    string `json:"state" bson:"state,omitempty"`
				District string `json:"district" bson:"district,omitempty"`
				Block    string `json:"block" bson:"block,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"blockCrop" bson:"blockCrop,omitempty"`
		StateLiveStock struct {
			StateLiveStock      string `json:"stateLiveStock" bson:"stateLiveStock,omitempty"`
			Commodity           string `json:"commodity" bson:"commodity,omitempty"`
			NameInLocalLanguage string `json:"nameInLocalLanguage" bson:"nameInLocalLanguage,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			Controls            string `json:"controls" bson:"controls,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Close               string `json:"close" bson:"close,omitempty"`
			Save                string `json:"save" bson:"save,omitempty"`
			NewStateLiveStock   string `json:"newStateLiveStock" bson:"newStateLiveStock,omitempty"`
			Filter              struct {
				State string `json:"state" bson:"state,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"stateLiveStock" bson:"stateLiveStock,omitempty"`
		BannedItems struct {
			BannedItems         string `json:"bannedItems" bson:"bannedItems,omitempty"`
			Name                string `json:"name" bson:"name,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Controls            string `json:"controls" bson:"controls,omitempty"`
			Close               string `json:"close" bson:"close,omitempty"`
			Save                string `json:"save" bson:"save,omitempty"`
			Desc                string `json:"description" bson:"description,omitempty"`
			Type                string `json:"type" bson:"type,omitempty"`
			NewBannedItems      string `json:"newBannedItems" bson:"newBannedItems,omitempty"`
			Pesticides          string `json:"pesticides" bson:"pesticides,omitempty"`
			Insecticide         string `json:"insecticide" bson:"insecticide,omitempty"`
			LivestockPainkiller string `json:livestockPainkiller" bson:"livestockPainkiller,omitempty"`
		} `json:"bannedItems" bson:"bannedItems,omitempty"`
		NarpZones struct {
			NarpZones    string `json:"narpZones" bson:"narpZones,omitempty"`
			Name         string `json:"name" bson:"name,omitempty"`
			Status       string `json:"status" bson:"status,omitempty"`
			Controls     string `json:"controls" bson:"controls,omitempty"`
			Close        string `json:"close" bson:"close,omitempty"`
			Save         string `json:"save" bson:"save,omitempty"`
			Zone         string `json:"zone" bson:"zone,omitempty"`
			Add          string `json:"add" bson:"add,omitempty"`
			NewNarpZones string `json:"newNarpZones" bson:"newNarpZones,omitempty"`
		} `json:"narpZones" bson:"narpZones,omitempty"`
		Vaccine struct {
			Vaccine    string `json:"vaccine" bson:"vaccine,omitempty"`
			Name       string `json:"name" bson:"name,omitempty"`
			Status     string `json:"status" bson:"status,omitempty"`
			Controls   string `json:"controls" bson:"controls,omitempty"`
			Close      string `json:"close" bson:"close,omitempty"`
			Save       string `json:"save" bson:"save,omitempty"`
			Desc       string `json:"description" bson:"description,omitempty"`
			NewVaccine string `json:"newVaccine" bson:"newVaccine,omitempty"`
		} `json:"vaccine" bson:"vaccine,omitempty"`
		LiveStockVaccine struct {
			LiveStockVaccine string `json:"liveStockVaccine" bson:"liveStockVaccine,omitempty"`
			Livestock        string `json:"livestock" bson:"livestock,omitempty"`
			Age              string `json:"age" bson:"age,omitempty"`
			Disease          string `json:"disease" bson:"disease,omitempty"`
			Immunity         string `json:"immunity" bson:"immunity,omitempty"`
			Dose             string `json:"dose" bson:"dose,omitempty"`
			Booster          string `json:"booster" bson:"booster,omitempty"`
			AnyTime          string `json:"anyTime" bson:"anyTime,omitempty"`
			Controls         string `json:"controls" bson:"controls,omitempty"`
			Status           string `json:"status" bson:"status,omitempty"`
			Close            string `json:"close" bson:"close,omitempty"`
			Save             string `json:"save" bson:"save,omitempty"`
			Vaccine          string `json:"vaccine" bson:"vaccine,omitempty"`
			PrintReport      string `json:"printReport" bson:"printReport,omitempty"`
			MonthFrom        string `json:"monthFrom" bson:"monthFrom,omitempty"`
			MonthTo          string `json:"monthTo" bson:"monthTo,omitempty"`
			BoosterTime      string `json:"boosterTime" bson:"boosterTime,omitempty"`
			BoosterDose      string `json:"boosterDose" bson:"boosterDose,omitempty"`
			Filter           struct {
				State     string `json:"state" bson:"state,omitempty"`
				Livestock string `json:"livestock" bson:"livestock,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"liveStockVaccine" bson:"liveStockVaccine,omitempty"`
		DistrictWeather struct {
			DistrictWeather    string `json:"districtWeather" bson:"districtWeather,omitempty"`
			SrNo               string `json:"srNo" bson:"srNo,omitempty"`
			From               string `json:"from" bson:"from,omitempty"`
			To                 string `json:"to" bson:"to,omitempty"`
			WindSpeedFrom      string `json:"windSpeedFrom" bson:"windSpeedFrom,omitempty"`
			WindSpeedTo        string `json:"windSpeedTo" bson:"windSpeedTo,omitempty"`
			RainfallFrom       string `json:"rainfallFrom" bson:"rainfallFrom,omitempty"`
			RainfallTo         string `json:"rainfallTo" bson:"rainfallTo,omitempty"`
			TemperatureMinFrom string `json:"temperatureMinFrom" bson:"temperatureMinFrom,omitempty"`
			TemperatureMinTo   string `json:"temperatureMinTo" bson:"temperatureMinTo,omitempty"`
			TemperatureMaxFrom string `json:"temperatureMaxFrom" bson:"temperatureMaxFrom,omitempty"`
			TemperatureMaxTo   string `json:"temperatureMaxTo" bson:"temperatureMaxTo,omitempty"`
			RhMinFrom          string `json:"rhMinFrom" bson:"rhMinFrom,omitempty"`
			RhMinTo            string `json:"rhMinTo" bson:"rhMinTo,omitempty"`
			RhMaxFrom          string `json:"rhMaxFrom" bson:"rhMaxFrom,omitempty"`
			RhMaxTo            string `json:"rhMaxTo" bson:"rhMaxTo,omitempty"`
			WindDirectionFrom  string `json:"windDirectionFrom" bson:"windDirectionFrom,omitempty"`
			WindDirectionTo    string `json:"windDirectionTo" bson:"windDirectionTo,omitempty"`
			NewDistrictWeather string `json:"newDistrictWeather" bson:"newDistrictWeather,omitempty"`
		} `json:"districtWeather" bson:"districtWeather,omitempty"`
		Organisation struct {
			Organisation          string `json:"organisation" bson:"organisation,omitempty"`
			Name                  string `json:"name" bson:"name,omitempty"`
			Status                string `json:"status" bson:"status,omitempty"`
			Controls              string `json:"controls" bson:"controls,omitempty"`
			Close                 string `json:"close" bson:"close,omitempty"`
			Save                  string `json:"save" bson:"save,omitempty"`
			Desc                  string `json:"description" bson:"description,omitempty"`
			NewOrganisation       string `json:"newOrganisation" bson:"newOrganisation,omitempty"`
			ValidationName        string `json:"validationName" bson:"validationName,omitempty"`
			ValidationDescription string `json:"validationDescription" bson:"validationDescription,omitempty"`
		} `json:"organisation" bson:"organisation,omitempty"`
		Project struct {
			Project       string `json:"project" bson:"project,omitempty"`
			ManageProject struct {
				Organisation    string `json:"organisation" bson:"organisation,omitempty"`
				ManageProject   string `json:"manageProject" bson:"manageProject,omitempty"`
				Name            string `json:"name" bson:"name,omitempty"`
				Status          string `json:"status" bson:"status,omitempty"`
				Controls        string `json:"controls" bson:"controls,omitempty"`
				Close           string `json:"close" bson:"close,omitempty"`
				Save            string `json:"save" bson:"save,omitempty"`
				Budget          string `json:"budget" bson:"budget,omitempty"`
				StartDate       string `json:"startDate" bson:"startDate,omitempty"`
				EndDate         string `json:"endDate" bson:"endDate,omitempty"`
				NewProject      string `json:"newProject" bson:"newProject,omitempty"`
				Email           string `json:"email" bson:"email,omitempty"`
				Patners         string `json:"patners" bson:"patners,omitempty"`
				State           string `json:"state" bson:"state,omitempty"`
				Remarks         string `json:"remarks" bson:"remarks,omitempty"`
				NationalLevel   string `json:"nationalLevel" bson:"nationalLevel,omitempty"`
				KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
				Sno             string `json:"sno" bson:"sno,omitempty"`
				FetchMembers    string `json:"fetchMembers" bson:"fetchMembers,omitempty"`
				CheckAll        string `json:"checkAll" bson:"checkAll,omitempty"`
				UnCheckAll      string `json:"unCheckAll" bson:"unCheckAll,omitempty"`
			} `json:"manageProject" bson:"manageProject,omitempty"`
		} `json:"project" bson:"project,omitempty"`
	} `json:"master" bson:"master,omitempty"`
	Users struct {
		Users           string `json:"users" bson:"users,omitempty"`
		SuperAdmin      string `json:"superAdmin" bson:"superAdmin,omitempty"`
		CallCenterAgent struct {
			CallCenterAgent       string `json:"callCenterAgent" bson:"callCenterAgent,omitempty"`
			FirstName             string `json:"firstName" bson:"firstName,omitempty"`
			LastName              string `json:"lastName" bson:"lastName,omitempty"`
			UserName              string `json:"userName" bson:"userName,omitempty"`
			Organisation          string `json:"organisation" bson:"organisation,omitempty"`
			Email                 string `json:"email" bson:"email,omitempty"`
			State                 string `json:"state" bson:"state,omitempty"`
			District              string `json:"district" bson:"district,omitempty"`
			MobileNumber          string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                string `json:"status" bson:"status,omitempty"`
			Active                string `json:"active" bson:"active,omitempty"`
			Disable               string `json:"disable" bson:"disable,omitempty"`
			CALLCENTERAGENTLIST   string `json:"callCenterAgentList" bson:"callCenterAgentList,omitempty"`
			SAVECALLCENTERAGENT   string `json:"saveCallCenterAgent" bson:"saveCallCenterAgent,omitempty"`
			ADDCALLCENTERAGENT    string `json:"addCallCentreAgent" bson:"addCallCentreAgent,omitempty"`
			UpdateCallCenterAgent string `json:"updateCallCenterAgent" bson:"updateCallCenterAgent,omitempty"`
			ViewCallCenterAgent   string `json:"viewCallCenterAgent" bson:"viewCallCenterAgent,omitempty"`
		} `json:"callCenterAgent" bson:"callCenterAgent,omitempty"`
		ContentCreator struct {
			ContentCreator       string `json:"contentCreator" bson:"contentCreator,omitempty"`
			ContentCreatorList   string `json:"contentCreatorList" bson:"contentCreatorList,omitempty"`
			FirstName            string `json:"firstName" bson:"firstName,omitempty"`
			LastName             string `json:"lastName" bson:"lastName,omitempty"`
			UserName             string `json:"userName" bson:"userName,omitempty"`
			Email                string `json:"email" bson:"email,omitempty"`
			State                string `json:"state" bson:"state,omitempty"`
			District             string `json:"district" bson:"district,omitempty"`
			MobileNumber         string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status               string `json:"status" bson:"status,omitempty"`
			Active               string `json:"active" bson:"active,omitempty"`
			Disable              string `json:"disable" bson:"disable,omitempty"`
			SaveContentCreator   string `json:"saveContentCreator" bson:"saveContentCreator,omitempty"`
			UpdateContentCreator string `json:"updateContentCreator" bson:"updateContentCreator,omitempty"`
			AddContentCreator    string `json:"addContentCreator" bson:"addContentCreator,omitempty"`
			ViewContentCreator   string `json:"viewContentCreator" bson:"viewContentCreator,omitempty"`
		} `json:"contentCreator" bson:"contentCreator,omitempty"`
		ContentManager struct {
			ContentManager       string `json:"contentManager" bson:"contentManager,omitempty"`
			UpdateContentManager string `json:"updateContentManager" bson:"updateContentManager,omitempty"`
			SaveContentManager   string `json:"saveContentManager" bson:"saveContentManager,omitempty"`
			AddContentManager    string `json:"addContentManager" bson:"addContentManager,omitempty"`
			ViewContentManager   string `json:"viewContentManager" bson:"viewContentManager,omitempty"`
			ContentManagerList   string `json:"contentManagerList" bson:"contentManagerList,omitempty"`
			FirstName            string `json:"firstName" bson:"firstName,omitempty"`
			LastName             string `json:"lastName" bson:"lastName,omitempty"`
			UserName             string `json:"userName" bson:"userName,omitempty"`
			Email                string `json:"email" bson:"email,omitempty"`
			State                string `json:"state" bson:"state,omitempty"`
			District             string `json:"district" bson:"district,omitempty"`
			MobileNumber         string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status               string `json:"status" bson:"status,omitempty"`
			Active               string `json:"active" bson:"active,omitempty"`
			Disable              string `json:"disable" bson:"disable,omitempty"`
		} `json:"contentManager" bson:"contentManager,omitempty"`
		ContentProvider struct {
			ContentProvider       string `json:"contentProvider" bson:"contentProvider,omitempty"`
			SaveContentProvider   string `json:"saveContentProvider" bson:"saveContentProvider,omitempty"`
			UpdateContentProvider string `json:"updateContentProvider" bson:"updateContentProvider,omitempty"`
			AddContentProvider    string `json:"addContentProvider" bson:"addContentProvider,omitempty"`
			ViewContentProvider   string `json:"viewContentProvider" bson:"viewContentProvider,omitempty"`
			CONTENTPROVIDELIST    string `json:"contentProviderList" bson:"contentProviderList,omitempty"`
			FirstName             string `json:"firstName" bson:"firstName,omitempty"`
			LastName              string `json:"lastName" bson:"lastName,omitempty"`
			UserName              string `json:"userName" bson:"userName,omitempty"`
			Email                 string `json:"email" bson:"email,omitempty"`
			State                 string `json:"state" bson:"state,omitempty"`
			District              string `json:"district" bson:"district,omitempty"`
			MobileNumber          string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                string `json:"status" bson:"status,omitempty"`
			Active                string `json:"active" bson:"active,omitempty"`
			Disable               string `json:"disable" bson:"disable,omitempty"`
		} `json:"contentProvider" bson:"contentProvider,omitempty"`
		ContentDisseminator struct {
			ContentDisseminator       string `json:"contentDisseminator" bson:"contentDisseminator,omitempty"`
			UpdateContentDisseminator string `json:"updateContentDisseminator" bson:"updateContentDisseminator,omitempty"`
			SaveContentDisseminator   string `json:"saveContentDisseminator" bson:"saveContentDisseminator,omitempty"`
			AddContentDisseminator    string `json:"addContentDisseminator" bson:"addContentDisseminator,omitempty"`
			ViewContentDisseminator   string `json:"viewContentDisseminator" bson:"viewContentDisseminator,omitempty"`
			DISSEMINATORLIST          string `json:"contentDisseminatorList" bson:"contentDisseminatorList,omitempty"`
			FirstName                 string `json:"firstName" bson:"firstName,omitempty"`
			LastName                  string `json:"lastName" bson:"lastName,omitempty"`
			UserName                  string `json:"userName" bson:"userName,omitempty"`
			Email                     string `json:"email" bson:"email,omitempty"`
			State                     string `json:"state" bson:"state,omitempty"`
			District                  string `json:"district" bson:"district,omitempty"`
			MobileNumber              string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                    string `json:"status" bson:"status,omitempty"`
			Active                    string `json:"active" bson:"active,omitempty"`
			Disable                   string `json:"disable" bson:"disable,omitempty"`
		} `json:"contentDisseminator" bson:"contentDisseminator,omitempty"`
		FieldAgent struct {
			FieldAgent       string `json:"fieldAgent" bson:"fieldAgent,omitempty"`
			SaveFieldAgent   string `json:"saveFieldAgent" bson:"saveFieldAgent,omitempty"`
			UpdateFieldAgent string `json:"updateFieldAgent" bson:"updateFieldAgent,omitempty"`
			ViewFieldAgent   string `json:"viewFieldAgent" bson:"viewFieldAgent,omitempty"`
			AddFieldAgent    string `json:"addFieldAgent" bson:"addFieldAgent,omitempty"`
			FIELDAGENTLIST   string `json:"fieldAgentList" bson:"fieldAgentList,omitempty"`
			FirstName        string `json:"firstName" bson:"firstName,omitempty"`
			LastName         string `json:"lastName" bson:"lastName,omitempty"`
			UserName         string `json:"userName" bson:"userName,omitempty"`
			Email            string `json:"email" bson:"email,omitempty"`
			State            string `json:"state" bson:"state,omitempty"`
			District         string `json:"district" bson:"district,omitempty"`
			MobileNumber     string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status           string `json:"status" bson:"status,omitempty"`
			Active           string `json:"active" bson:"active,omitempty"`
			Disable          string `json:"disable" bson:"disable,omitempty"`
		} `json:"fieldAgent" bson:"fieldAgent,omitempty"`
		LanguageTranslator struct {
			LanguageTranslator       string `json:"languageTranslator" bson:"languageTranslator,omitempty"`
			SaveLanguageTranslator   string `json:"saveLanguageTranslator" bson:"saveLanguageTranslator,omitempty"`
			UpdateLanguageTranslator string `json:"updateLanguageTranslator" bson:"updateLanguageTranslator,omitempty"`
			ViewLanguageTranslator   string `json:"viewLanguageTranslator" bson:"viewLanguageTranslator,omitempty"`
			AddLanguageTranslator    string `json:"addLanguageTranslator" bson:"addLanguageTranslator,omitempty"`
			LANGUAGETRANSLATORLIST   string `json:"languageTranslatorList" bson:"languageTranslatorList,omitempty"`
			FirstName                string `json:"firstName" bson:"firstName,omitempty"`
			LastName                 string `json:"lastName" bson:"lastName,omitempty"`
			UserName                 string `json:"userName" bson:"userName,omitempty"`
			Email                    string `json:"email" bson:"email,omitempty"`
			State                    string `json:"state" bson:"state,omitempty"`
			District                 string `json:"district" bson:"district,omitempty"`
			MobileNumber             string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                   string `json:"status" bson:"status,omitempty"`
			Active                   string `json:"active" bson:"active,omitempty"`
			Disable                  string `json:"disable" bson:"disable,omitempty"`
		} `json:"languageTranslator" bson:"languageTranslator,omitempty"`
		TranslationApprover struct {
			TranslationApprover       string `json:"translationApprover" bson:"translationApprover,omitempty"`
			SaveTranslationApprover   string `json:"saveTranslationApprover" bson:"saveTranslationApprover,omitempty"`
			UpdateTranslationApprover string `json:"updateTranslationApprover" bson:"updateTranslationApprover,omitempty"`
			ViewTranslationApprover   string `json:"viewTranslationApprover" bson:"viewTranslationApprover,omitempty"`
			AddTranslationApprover    string `json:"addTranslationApprover" bson:"addTranslationApprover,omitempty"`
			TRANSLATIONAPPROVERLIST   string `json:"translationApproverList" bson:"translationApproverList,omitempty"`
			FirstName                 string `json:"firstName" bson:"firstName,omitempty"`
			LastName                  string `json:"lastName" bson:"lastName,omitempty"`
			UserName                  string `json:"userName" bson:"userName,omitempty"`
			Email                     string `json:"email" bson:"email,omitempty"`
			State                     string `json:"state" bson:"state,omitempty"`
			District                  string `json:"district" bson:"district,omitempty"`
			MobileNumber              string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                    string `json:"status" bson:"status,omitempty"`
			Active                    string `json:"active" bson:"active,omitempty"`
			Disable                   string `json:"disable" bson:"disable,omitempty"`
		} `json:"translationApprover" bson:"translationApprover,omitempty"`
		Management struct {
			Management       string `json:"management" bson:"management,omitempty"`
			SaveManagement   string `json:"saveManagement" bson:"saveManagement,omitempty"`
			UpdateManagement string `json:"updateManagement" bson:"updateManagement,omitempty"`
			ViewManagement   string `json:"viewManagement" bson:"viewManagement,omitempty"`
			AddManagement    string `json:"addManagement" bson:"addManagement,omitempty"`
			FirstName        string `json:"firstName" bson:"firstName,omitempty"`
			LastName         string `json:"lastName" bson:"lastName,omitempty"`
			UserName         string `json:"userName" bson:"userName,omitempty"`
			Organisation     string `json:"organisation" bson:"organisation,omitempty"`
			Email            string `json:"email" bson:"email,omitempty"`
			State            string `json:"state" bson:"state,omitempty"`
			District         string `json:"district" bson:"district,omitempty"`
			MobileNumber     string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status           string `json:"status" bson:"status,omitempty"`
			Active           string `json:"active" bson:"active,omitempty"`
			Disable          string `json:"disable" bson:"disable,omitempty"`
		} `json:"management" bson:"management,omitempty"`
		DistrictAdmin struct {
			DistrictAdmin       string `json:"districtAdmin" bson:"districtAdmin,omitempty"`
			SaveDistrictAdmin   string `json:"saveDistrictAdmin" bson:"saveDistrictAdmin,omitempty"`
			UpdateDistrictAdmin string `json:"updateDistrictAdmin" bson:"updateDistrictAdmin,omitempty"`
			ViewDistrictAdmin   string `json:"viewDistrictAdmin" bson:"viewDistrictAdmin,omitempty"`
			AddDistrictAdmin    string `json:"addDistrictAdmin" bson:"addDistrictAdmin,omitempty"`
			FirstName           string `json:"firstName" bson:"firstName,omitempty"`
			LastName            string `json:"lastName" bson:"lastName,omitempty"`
			UserName            string `json:"userName" bson:"userName,omitempty"`
			Email               string `json:"email" bson:"email,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			District            string `json:"district" bson:"district,omitempty"`
			MobileNumber        string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Active              string `json:"active" bson:"active,omitempty"`
			Disable             string `json:"disable" bson:"disable,omitempty"`
		} `json:"districtAdmin" bson:"districtAdmin,omitempty"`
		SubjectMatteExpert struct {
			SubjectMatteExpert       string `json:"subjectMatteExpert" bson:"subjectMatteExpert,omitempty"`
			SubjectMatterExpertList  string `json:"subjectMatteExpertList" bson:"subjectMatteExpertList,omitempty"`
			SaveSubjectMatteExpert   string `json:"saveSubjectMatteExpert" bson:"saveSubjectMatteExpert,omitempty"`
			ViewSubjectMatteExpert   string `json:"viewSubjectMatteExpert" bson:"viewSubjectMatteExpert,omitempty"`
			UpdateSubjectMatteExpert string `json:"updateSubjectMatteExpert" bson:"updateSubjectMatteExpert,omitempty"`
			AddSubjectMatteExpert    string `json:"addSubjectMatteExpert" bson:"addSubjectMatteExpert,omitempty"`
			FirstName                string `json:"firstName" bson:"firstName,omitempty"`
			LastName                 string `json:"lastName" bson:"lastName,omitempty"`
			UserName                 string `json:"userName" bson:"userName,omitempty"`
			Organisation             string `json:"organisation" bson:"organisation,omitempty"`
			Email                    string `json:"email" bson:"email,omitempty"`
			State                    string `json:"state" bson:"state,omitempty"`
			District                 string `json:"district" bson:"district,omitempty"`
			MobileNumber             string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                   string `json:"status" bson:"status,omitempty"`
			Active                   string `json:"active" bson:"active,omitempty"`
			Disable                  string `json:"disable" bson:"disable,omitempty"`
		} `json:"subjectMatteExpert" bson:"subjectMatteExpert,omitempty"`
		SystemAdmin struct {
			SystemAdmin       string `json:"systemAdmin" bson:"systemAdmin,omitempty"`
			SaveSystemAdmin   string `json:"saveSystemAdmin" bson:"saveSystemAdmin,omitempty"`
			UpdateSystemAdmin string `json:"updateSystemAdmin" bson:"updateSystemAdmin,omitempty"`
			ViewSystemAdmin   string `json:"viewSystemAdmin" bson:"viewSystemAdmin,omitempty"`
			AddSystemAdmin    string `json:"addSystemAdmin" bson:"addSystemAdmin,omitempty"`
			FirstName         string `json:"firstName" bson:"firstName,omitempty"`
			LastName          string `json:"lastName" bson:"lastName,omitempty"`
			UserName          string `json:"userName" bson:"userName,omitempty"`
			Email             string `json:"email" bson:"email,omitempty"`
			State             string `json:"state" bson:"state,omitempty"`
			District          string `json:"district" bson:"district,omitempty"`
			MobileNumber      string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status            string `json:"status" bson:"status,omitempty"`
			Active            string `json:"active" bson:"active,omitempty"`
			Disable           string `json:"disable" bson:"disable,omitempty"`
		} `json:"systemAdmin" bson:"systemAdmin,omitempty"`
		VisitorViewer struct {
			VisitorViewer       string `json:"visitorViewer" bson:"visitorViewer,omitempty"`
			SaveVisitorViewer   string `json:"saveVisitorViewer" bson:"saveVisitorViewer,omitempty"`
			UpdateVisitorViewer string `json:"updateVisitorViewer" bson:"updateVisitorViewer,omitempty"`
			AddVisitorViewer    string `json:"addVisitorViewer" bson:"addVisitorViewer,omitempty"`
			ViewVisitorViewer   string `json:"viewVisitorViewer" bson:"viewVisitorViewer,omitempty"`
			VisitorViewers      string `json:"visitorViewers" bson:"visitorViewers,omitempty"`
			FirstName           string `json:"firstName" bson:"firstName,omitempty"`
			LastName            string `json:"lastName" bson:"lastName,omitempty"`
			UserName            string `json:"userName" bson:"userName,omitempty"`
			Role                string `json:"role" bson:"role,omitempty"`
			Email               string `json:"email" bson:"email,omitempty"`
			MobileNumber        string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Active              string `json:"active" bson:"active,omitempty"`
			Disable             string `json:"disable" bson:"disable,omitempty"`
		} `json:"visitorViewer" bson:"visitorViewer,omitempty"`
		AllUsers struct {
			AllUsers     string `json:"allUsers" bson:"allUsers,omitempty"`
			FirstName    string `json:"firstName" bson:"firstName,omitempty"`
			LastName     string `json:"lastName" bson:"lastName,omitempty"`
			UserName     string `json:"userName" bson:"userName,omitempty"`
			Role         string `json:"role" bson:"role,omitempty"`
			Email        string `json:"email" bson:"email,omitempty"`
			MobileNumber string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status       string `json:"status" bson:"status,omitempty"`
			Active       string `json:"active" bson:"active,omitempty"`
			Disable      string `json:"disable" bson:"disable,omitempty"`
		} `json:"allUsers" bson:"allUsers,omitempty"`
		Trainer struct {
			Trainer       string `json:"trainer" bson:"trainer,omitempty"`
			SaveTrainer   string `json:"saveTrainer" bson:"saveTrainer,omitempty"`
			UpdateTrainer string `json:"updateTrainer" bson:"updateTrainer,omitempty"`
			AddTrainer    string `json:"addTrainer" bson:"addTrainer,omitempty"`
			ViewTrainer   string `json:"viewTrainer" bson:"viewTrainer,omitempty"`
			FirstName     string `json:"firstName" bson:"firstName,omitempty"`
			LastName      string `json:"lastName" bson:"lastName,omitempty"`
			UserName      string `json:"userName" bson:"userName,omitempty"`
			Email         string `json:"email" bson:"email,omitempty"`
			MobileNumber  string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status        string `json:"status" bson:"status,omitempty"`
			Active        string `json:"active" bson:"active,omitempty"`
			Disable       string `json:"disable" bson:"disable,omitempty"`
		} `json:"trainer" bson:"trainer,omitempty"`
		GuestUsers struct {
			GuestUsers       string `json:"guestUsers" bson:"guestUsers,omitempty"`
			SaveGuestUsers   string `json:"saveGuestUsers" bson:"saveGuestUsers,omitempty"`
			UpdateGuestUsers string `json:"updateGuestUsers" bson:"updateGuestUsers,omitempty"`
			AddGuestUsers    string `json:"addGuestUsers" bson:"addGuestUsers,omitempty"`
			ViewGuestUsers   string `json:"viewGuestUsers" bson:"viewGuestUsers,omitempty"`
			UserName         string `json:"userName" bson:"userName,omitempty"`
			Email            string `json:"email" bson:"email,omitempty"`
			MobileNumber     string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status           string `json:"status" bson:"status,omitempty"`
			Active           string `json:"active" bson:"active,omitempty"`
			Disable          string `json:"disable" bson:"disable,omitempty"`
		} `json:"guestUsers" bson:"guestUsers,omitempty"`
		FieldAgentLead struct {
			FieldAgentLead       string `json:"fieldAgentLead" bson:"fieldAgentLead,omitempty"`
			SaveFieldAgentLead   string `json:"saveFieldAgentLead" bson:"saveFieldAgentLead,omitempty"`
			UpdateFieldAgentLead string `json:"updateFieldAgentLead" bson:"updateFieldAgentLead,omitempty"`
			AddFieldAgentLead    string `json:"addFieldAgentLead" bson:"addFieldAgentLead,omitempty"`
			ViewFieldAgentLead   string `json:"viewFieldAgentLead" bson:"viewFieldAgentLead,omitempty"`
			FirstName            string `json:"firstName" bson:"firstName,omitempty"`
			LastName             string `json:"lastName" bson:"lastName,omitempty"`
			UserName             string `json:"userName" bson:"userName,omitempty"`
			Email                string `json:"email" bson:"email,omitempty"`
			State                string `json:"state" bson:"state,omitempty"`
			District             string `json:"district" bson:"district,omitempty"`
			MobileNumber         string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status               string `json:"status" bson:"status,omitempty"`
			Active               string `json:"active" bson:"active,omitempty"`
			Disable              string `json:"disable" bson:"disable,omitempty"`
		} `json:"fieldAgentLead" bson:"fieldAgentLead,omitempty"`
		SelfRegistration struct {
			SelfRegistration               string `json:"selfRegistration" bson:"selfRegistration,omitempty"`
			Approved                       string `json:"approved" bson:"approved,omitempty"`
			Rejected                       string `json:"rejected" bson:"rejected,omitempty"`
			Requested                      string `json:"requested" bson:"requested,omitempty"`
			FirstName                      string `json:"firstName" bson:"firstName,omitempty"`
			LastName                       string `json:"lastName" bson:"lastName,omitempty"`
			UserName                       string `json:"userName" bson:"userName,omitempty"`
			Role                           string `json:"role" bson:"role,omitempty"`
			Email                          string `json:"email" bson:"email,omitempty"`
			MobileNumber                   string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Status                         string `json:"status" bson:"status,omitempty"`
			APPROVEREJECTSELFREGISTRATIONS string `json:"approveReject" bson:"approveReject,omitempty"`
		} `json:"selfRegistration" bson:"selfRegistration,omitempty"`
		CreateUser struct {
			Organisation                 string `json:"organisation" bson:"organisation,omitempty"`
			Name                         string `json:"name" bson:"name,omitempty"`
			UserName                     string `json:"userName" bson:"userName,omitempty"`
			FirstName                    string `json:"firstName" bson:"firstName,omitempty"`
			Lastname                     string `json:"lastname" bson:"lastname,omitempty"`
			Project                      string `json:"project" bson:"project,omitempty"`
			AlternateNumber              string `json:"alternateNumber" bson:"alternateNumber,omitempty"`
			FatherName                   string `json:"fatherName" bson:"fatherName,omitempty"`
			Father_HusbandName           string `json:"father_husbandName" bson:"father_husbandName,omitempty"`
			State                        string `json:"state" bson:"state,omitempty"`
			District                     string `json:"district" bson:"district,omitempty"`
			Block                        string `json:"block" bson:"block,omitempty"`
			Grampanchayat                string `json:"grampanchayat" bson:"grampanchayat,omitempty"`
			City                         string `json:"city" bson:"city,omitempty"`
			PinCode                      string `json:"pinCode" bson:"pinCode,omitempty"`
			Gender                       string `json:"gender" bson:"gender,omitempty"`
			MobileNumber                 string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Email                        string `json:"email" bson:"email,omitempty"`
			DateOfBirth                  string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
			Address                      string `json:"address" bson:"address,omitempty"`
			AccessLevel                  string `json:"accessLevel" bson:"accessLevel,omitempty"`
			MailNotification             string `json:"mailNotification" bson:"mailNotification"`
			SmsNotification              string `json:"smsNotification" bson:"smsNotification,omitempty"`
			ViewLanguage                 string `json:"viewLanguage" bson:"viewLanguage"`
			EducationalQualification     string `json:"educationalQualification" bson:"educationalQualification,omitempty"`
			KnowledgeDomains             string `json:"knowledgeDomains" bson:"knowledgeDomains,omitempty"`
			SubDomains                   string `json:"subDomains" bson:"subDomains,omitempty"`
			Experience                   string `json:"experience" bson:"experience,omitempty"`
			LanguageExperience           string `json:"languageExperience" bson:"languageExperience"`
			SmsCount_Day                 string `json:"smsCount_day" bson:"smsCount_day,omitempty"`
			KdExpertise                  string `json:"kdExpertise" bson:"kdExpertise,omitempty"`
			OrganisationNatureofBusiness string `json:"organisationNatureofBusiness" bson:"organisationNatureofBusiness,omitempty"`
			Officenumber                 string `json:"officenumber" bson:"officenumber,omitempty"`
			Officeaddress                string `json:"officeaddress" bson:"officeaddress,omitempty"`
			Designation                  string `json:"designation" bson:"designation,omitempty"`
			LanguagesKnown               string `json:"languagesKnown" bson:"languagesKnown"`
			Village                      string `json:"village" bson:"village,omitempty"`
			Occupation                   string `json:"occupation" bson:"occupation,omitempty"`
			UserType                     string `json:"userType" bson:"userType,omitempty"`
			ProofType                    string `json:"proofType" bson:"proofType,omitempty"`
			ProofNo                      string `json:"proofNo" bson:"proofNo,omitempty"`
			SubjectExpertise             string `json:"subjectExpertise" bson:"subjectExpertise,omitempty"`
			BloodGroup                   string `json:"bloodGroup" bson:"bloodGroup,omitempty"`
			SelectGender                 string `json:"selectGender" bson:"selectGender,omitempty"`
			Male                         string `json:"male" bson:"male,omitempty"`
			FeMale                       string `json:"feMale" bson:"feMale,omitempty"`
			Others                       string `json:"others" bson:"others,omitempty"`
			MobileNumberAlert            string `json:"mobileNumberAlert" bson:"mobileNumberAlert,omitempty"`
			AddedSuccessfully            string `json:"addedSuccessfully" bson:"addedSuccessfully,omitempty"`
			UpdatedSuccessfully          string `json:"updatedSuccessfully" bson:"updatedSuccessfully,omitempty"`
			ApprovedSuccess              string `json:"approvedSuccess" bson:"approvedSuccess,omitempty"`
		} `json:"createUser" bson:"createUser,omitempty"`
		Control struct {
			Role struct {
				Role  string `json:"role" bson:"role,omitempty"`
				Close string `json:"close" bson:"close,omitempty"`
				Save  string `json:"save" bson:"save,omitempty"`
			} `json:"role" bson:"role,omitempty"`
			Password struct {
				Password        string `json:"password" bson:"password,omitempty"`
				NewPassword     string `json:"newPassword" bson:"newPassword,omitempty"`
				ConfirmPassword string `json:"confirmPassword" bson:"confirmPassword,omitempty"`
				Close           string `json:"close" bson:"close,omitempty"`
				Save            string `json:"save" bson:"save,omitempty"`
			} `json:"password" bson:"password,omitempty"`
			Delete struct {
				Delete      string `json:"delete" bson:"delete,omitempty"`
				DeleteAlert string `json:"deleteAlert" bson:"deleteAlert,omitempty"`
				OK          string `json:"ok" bson:"ok,omitempty"`
				Cancel      string `json:"cancel" bson:"cancel,omitempty"`
			} `json:"delete" bson:"delete,omitempty"`
			Features struct {
				UserSettings struct {
					UserSettings string `json:"userSettings" bson:"userSettings,omitempty"`
					User         struct {
						User  string `json:"user" bson:"user,omitempty"`
						Read  string `json:"read" bson:"read,omitempty"`
						Write string `json:"write" bson:"write,omitempty"`
					} `json:"user" bson:"user,omitempty"`
					Location struct {
						Location string `json:"location" bson:"location,omitempty"`
						Read     string `json:"read" bson:"read,omitempty"`
						Write    string `json:"write" bson:"write,omitempty"`
					} `json:"location" bson:"location,omitempty"`
					KnowledgeDomain struct {
						KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
						Read            string `json:"read" bson:"read,omitempty"`
						Write           string `json:"write" bson:"write,omitempty"`
					} `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
					Crop struct {
						Crop  string `json:"crop" bson:"crop,omitempty"`
						Read  string `json:"read" bson:"read,omitempty"`
						Write string `json:"write" bson:"write,omitempty"`
					} `json:"crop" bson:"crop,omitempty"`
					Livestock struct {
						Livestock string `json:"livestock" bson:"livestock,omitempty"`
						Read      string `json:"read" bson:"read,omitempty"`
						Write     string `json:"write" bson:"write,omitempty"`
					} `json:"livestock" bson:"livestock,omitempty"`
					StateSeason struct {
						StateSeason string `json:"stateSeason" bson:"stateSeason,omitempty"`
						Read        string `json:"read" bson:"read,omitempty"`
						Write       string `json:"write" bson:"write,omitempty"`
					} `json:"stateSeason" bson:"stateSeason,omitempty"`
					StateCalendar struct {
						StateCalendar string `json:"stateCalendar" bson:"stateCalendar,omitempty"`
						Read          string `json:"read" bson:"read,omitempty"`
						Write         string `json:"write" bson:"write,omitempty"`
					} `json:"stateCalendar" bson:"stateCalendar,omitempty"`
					Farmer struct {
						Farmer string `json:"farmer" bson:"farmer,omitempty"`
						Read   string `json:"read" bson:"read,omitempty"`
						Write  string `json:"write" bson:"write,omitempty"`
					} `json:"farmer" bson:"farmer,omitempty"`
					ContingencyPlan struct {
						ContingencyPlan string `json:"contingencyPlan" bson:"contingencyPlan,omitempty"`
						Read            string `json:"read" bson:"read,omitempty"`
						Write           string `json:"write" bson:"write,omitempty"`
					} `json:"contingencyPlan" bson:"contingencyPlan,omitempty"`
					Market struct {
						Market string `json:"market" bson:"market,omitempty"`
						Read   string `json:"read" bson:"read,omitempty"`
						Write  string `json:"write" bson:"write,omitempty"`
					} `json:"market" bson:"market,omitempty"`
					Language struct {
						Language string `json:"language" bson:"language,omitempty"`
						Read     string `json:"read" bson:"read,omitempty"`
						Write    string `json:"write" bson:"write,omitempty"`
					} `json:"language" bson:"language,omitempty"`
					SoilType struct {
						SoilType string `json:"soilType" bson:"soilType,omitempty"`
						Read     string `json:"read" bson:"read,omitempty"`
						Write    string `json:"write" bson:"write,omitempty"`
					} `json:"soilType" bson:"soilType,omitempty"`
					Asset struct {
						Asset string `json:"asset" bson:"asset,omitempty"`
						Read  string `json:"read" bson:"read,omitempty"`
						Write string `json:"write" bson:"write,omitempty"`
					} `json:"asset" bson:"asset,omitempty"`
					Insect struct {
						Insect string `json:"insect" bson:"insect,omitempty"`
						Read   string `json:"read" bson:"read,omitempty"`
						Write  string `json:"write" bson:"write,omitempty"`
					} `json:"insect" bson:"insect,omitempty"`
					Disease struct {
						Disease string `json:"disease" bson:"disease,omitempty"`
						Read    string `json:"read" bson:"read,omitempty"`
						Write   string `json:"write" bson:"write,omitempty"`
					} `json:"disease" bson:"disease,omitempty"`
					BannedItem struct {
						BannedItem string `json:"bannedItem" bson:"bannedItem,omitempty"`
						Read       string `json:"read" bson:"read,omitempty"`
						Write      string `json:"write" bson:"write,omitempty"`
					} `json:"bannedItem" bson:"bannedItem,omitempty"`
					NARPZones struct {
						NarpZones string `json:"narpZones" bson:"narpZones,omitempty"`
						Read      string `json:"read" bson:"read,omitempty"`
						Write     string `json:"write" bson:"write,omitempty"`
					} `json:"narpZones" bson:"narpZones,omitempty"`
					Vaccines struct {
						Vaccines string `json:"vaccines" bson:"vaccines,omitempty"`
						Read     string `json:"read" bson:"read,omitempty"`
						Write    string `json:"write" bson:"write,omitempty"`
					} `json:"vaccines" bson:"vaccines,omitempty"`
					LivestockVaccines struct {
						LivestockVaccines string `json:"livestockVaccines" bson:"livestockVaccines,omitempty"`
						Read              string `json:"read" bson:"read,omitempty"`
						Write             string `json:"write" bson:"write,omitempty"`
					} `json:"livestockVaccines" bson:"livestockVaccines,omitempty"`
					DistrictWeather struct {
						DistrictWeather string `json:"districtWeather" bson:"districtWeather,omitempty"`
						Read            string `json:"read" bson:"read,omitempty"`
						Write           string `json:"write" bson:"write,omitempty"`
					} `json:"districtWeather" bson:"districtWeather,omitempty"`
					Organisation struct {
						Organisation string `json:"organisation" bson:"organisation,omitempty"`
						Read         string `json:"read" bson:"read,omitempty"`
						Write        string `json:"write" bson:"write,omitempty"`
					} `json:"organisation" bson:"organisation,omitempty"`
					Project struct {
						Project string `json:"project" bson:"project,omitempty"`
						Read    string `json:"read" bson:"read,omitempty"`
						Write   string `json:"write" bson:"write,omitempty"`
					} `json:"project" bson:"project,omitempty"`
					AidLocation struct {
						AidLocation string `json:"aidLocation" bson:"aidLocation,omitempty"`
						Read        string `json:"read" bson:"read,omitempty"`
						Write       string `json:"write" bson:"write,omitempty"`
					} `json:"aidLocation" bson:"aidLocation,omitempty"`
				} `json:"userSettings" bson:"userSettings,omitempty"`
				ContentAndQueryFeatures struct {
					ContentAndQueryFeatures string `json:"contentAndQueryFeatures" bson:"contentAndQueryFeatures,omitempty"`
					Action                  string `json:"action" bson:"action,omitempty"`
					CreateEdit              string `json:"createEdit" bson:"createEdit,omitempty"`
					BypassContent           string `json:"bypassContent" bson:"bypassContent,omitempty"`
					Manage                  string `json:"manage" bson:"manage,omitempty"`
					TranslateReview         string `json:"translateReview" bson:"translateReview,omitempty"`
					Upload                  string `json:"upload" bson:"upload,omitempty"`
					QueryEdit               string `json:"queryEdit" bson:"queryEdit,omitempty"`
					Delete                  string `json:"delete" bson:"delete,omitempty"`
					Review                  string `json:"review" bson:"review,omitempty"`
					Translate               string `json:"translate" bson:"translate,omitempty"`
					Disseminate             string `json:"disseminate" bson:"disseminate,omitempty"`
					Search                  string `json:"search" bson:"search,omitempty"`
				} `json:"contentAndQueryFeatures" bson:"contentAndQueryFeatures,omitempty"`
				SpecialFeatures struct {
					SpecialFeatures  string `json:"specialFeatures" bson:"specialFeatures,omitempty"`
					Action           string `json:"action" bson:"action,omitempty"`
					WeatherData      string `json:"weatherData" bson:"weatherData,omitempty"`
					PickResolveQuery string `json:"pickResolveQuery" bson:"pickResolveQuery,omitempty"`
					ViewOnlineUsers  string `json:"viewOnlineUsers" bson:"viewOnlineUsers,omitempty"`
				} `json:"specialFeatures" bson:"specialFeatures,omitempty"`
				Features string `json:"features" bson:"features,omitempty"`
				Close    string `json:"close" bson:"close,omitempty"`
				Add      string `json:"add" bson:"add,omitempty"`
			} `json:"features" bson:"features,omitempty"`
		} `json:"control" bson:"control,omitempty"`
	} `json:"users" bson:"users,omitempty"`
	Farmers struct {
		Farmers             string `json:"farmers" bson:"farmers,omitempty"`
		DynamicRegistration struct {
			DynamicRegistration string `json:"dynamicRegistration" bson:"dynamicRegistration,omitempty"`
			Name                string `json:"name" bson:"name,omitempty"`
			Gender              string `json:"gender" bson:"gender,omitempty"`
			SelectGender        string `json:"selectGender" bson:"selectGender,omitempty"`
			FatherName          string `json:"fatherName" bson:"fatherName,omitempty"`
			Mobilenumber        string `json:"mobilenumber" bson:"mobilenumber,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			District            string `json:"district" bson:"district,omitempty"`
			Block               string `json:"block" bson:"block,omitempty"`
			Grampanchyat        string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
			Village             string `json:"village" bson:"village,omitempty"`
			Add                 string `json:"add" bson:"add,omitempty"`
			Close               string `json:"close" bson:"close,omitempty"`
			Save                string `json:"save" bson:"save,omitempty"`
			Reset               string `json:"reset" bson:"reset,omitempty"`
		} `json:"dynamicRegistration" bson:"dynamicRegistration,omitempty"`
		DetailsRegistration struct {
			DetailsRegistration                 string `json:"detailsRegistration" bson:"detailsRegistration,omitempty"`
			Organisation                        string `json:"organisation" bson:"organisation,omitempty"`
			Name                                string `json:"name" bson:"name,omitempty"`
			DateOfBirth                         string `json:"dateOfBirth" bson:"dateOfBirth,omitempty"`
			Gender                              string `json:"gender" bson:"gender,omitempty"`
			FatherName                          string `json:"fatherName" bson:"fatherName,omitempty"`
			SpouseName                          string `json:"spouseName" bson:"spouseName,omitempty"`
			Education                           string `json:"education" bson:"education,omitempty"`
			YearlyIncome                        string `json:"yearlyIncome" bson:"yearlyIncome,omitempty"`
			AadhaarNumber                       string `json:"aadhaarNumber" bson:"aadhaarNumber,omitempty"`
			DoorNo                              string `json:"doorNo" bson:"doorNo,omitempty"`
			Street                              string `json:"street" bson:"street,omitempty"`
			LandMark                            string `json:"landMark" bson:"landMark,omitempty"`
			State                               string `json:"state" bson:"state,omitempty"`
			District                            string `json:"district" bson:"district,omitempty"`
			Block                               string `json:"block" bson:"block,omitempty"`
			Grampanchyat                        string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
			Village                             string `json:"village" bson:"village,omitempty"`
			Landline                            string `json:"landline" bson:"landline,omitempty"`
			MobileNumber                        string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			AlternateNumber                     string `json:"AlternateNumber" bson:"AlternateNumber,omitempty"`
			IsLeadWomenHasMobileAndNotEnteredYe string `json:"IsLeadWomenHasMobileAndNotEnteredYe" bson:"IsLeadWomenHasMobileAndNotEnteredYe,omitempty"`
			HasKitchenGarden                    string `json:"HasKitchenGarden" bson:"HasKitchenGarden,omitempty"`
			MemberOfMGNREGA                     string `json:"MemberOfMGNREGA" bson:"MemberOfMGNREGA,omitempty"`
			HasCredits                          string `json:"HasCredits" bson:"HasCredits,omitempty"`
			IsPhysicallyDisabled                string `json:"IsPhysicallyDisabled" bson:"IsPhysicallyDisabled,omitempty"`
			IsAnyMemberOfFamilyInCBO            string `json:"IsAnyMemberOfFamilyInCBO" bson:"IsAnyMemberOfFamilyInCBO,omitempty"`
			Asset                               string `json:"asset" bson:"asset,omitempty"`
			Caste                               string `json:"caste" bson:"caste,omitempty"`
			LikeToReceiveSMS                    string `json:"LikeToReceiveSMS" bson:"LikeToReceiveSMS,omitempty"`
			LikeToReceiveVoiceCall              string `json:"LikeToReceiveVoiceCall" bson:"LikeToReceiveVoiceCall,omitempty"`
			PreferredMarket                     string `json:"PreferredMarket" bson:"PreferredMarket,omitempty"`
			Previous                            string `json:"previous" bson:"previous,omitempty"`
			NextStep                            string `json:"nextStep" bson:"nextStep,omitempty"`
			Add                                 string `json:"add" bson:"add,omitempty"`
			Close                               string `json:"close" bson:"close,omitempty"`
			Save                                string `json:"save" bson:"save,omitempty"`
			Reset                               string `json:"reset" bson:"reset,omitempty"`
		} `json:"detailsRegistration" bson:"detailsRegistration,omitempty"`
		FarmerList struct {
			FarmerList string `json:"farmerList" bson:"farmerList,omitempty"`
			List       struct {
				Name           string `json:"name" bson:"name,omitempty"`
				FarmerId       string `json:"farmerId" bson:"farmerId,omitempty"`
				MobileNumber   string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
				State          string `json:"state" bson:"state,omitempty"`
				District       string `json:"district" bson:"district,omitempty"`
				Organisation   string `json:"organisation" bson:"organisation,omitempty"`
				Projects       string `json:"projects" bson:"projects,omitempty"`
				TotalLand      string `json:"totalLand" bson:"totalLand,omitempty"`
				CultivatedArea string `json:"cultivatedArea" bson:"cultivatedArea,omitempty"`
				VacantArea     string `json:"vacantArea" bson:"vacantArea,omitempty"`
				NoOfLivestock  string `json:"noOfLivestock" bson:"noOfLivestock,omitempty"`
				Caste          string `json:"caste" bson:"caste,omitempty"`
				Status         string `json:"status" bson:"status,omitempty"`
				Control        string `json:"control" bson:"control,omitempty"`
			} `json:"list" bson:"list,omitempty"`
			FeedBack struct {
				Feedback string `json:"feedback" bson:"feedback,omitempty"`
				Content  string `json:"content" bson:"content,omitempty"`
				Date     string `json:"date" bson:"date,omitempty"`
				Rating   string `json:"rating" bson:"rating,omitempty"`
				Type     string `json:"type" bson:"type,omitempty"`
			} `json:"feedBack" bson:"feedBack,omitempty"`
			Query struct {
				Query string `json:"query" bson:"query,omitempty"`
				Date  string `json:"date" bson:"date,omitempty"`
				Image string `json:"image" bson:"image,omitempty"`
			} `json:"query" bson:"query,omitempty"`
			Land struct {
				KhasraNo       string `json:"khasraNo" bson:"khasraNo,omitempty"`
				ParcelNo       string `json:"parcelNo" bson:"parcelNo,omitempty"`
				SoilType       string `json:"soilType" bson:"soilType,omitempty"`
				Ownership      string `json:"ownership" bson:"ownership,omitempty"`
				Area           string `json:"area" bson:"area,omitempty"`
				CultivatedArea string `json:"cultivatedArea" bson:"cultivatedArea,omitempty"`
				VacentArea     string `json:"vacentArea" bson:"vacentArea,omitempty"`
				Crop           string `json:"crop" bson:"crop,omitempty"`
				Controls       string `json:"controls" bson:"controls,omitempty"`
			} `json:"land" bson:"land,omitempty"`
			AddLand struct {
				PlotNo              string `json:"plotNo" bson:"plotNo,omitempty"`
				State               string `json:"state" bson:"state,omitempty"`
				District            string `json:"district" bson:"district,omitempty"`
				Block               string `json:"block" bson:"block,omitempty"`
				Grampanchyat        string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village             string `json:"village" bson:"village,omitempty"`
				Farmer              string `json:"farmer" bson:"farmer,omitempty"`
				IrrigatedType       string `json:"irrigateType" bson:"irrigateType,omitempty"`
				CultivatedArea      string `json:"cultivatedArea" bson:"cultivatedArea,omitempty"`
				VacentArea          string `json:"vacentArea" bson:"vacentArea,omitempty"`
				Area                string `json:"area" bson:"area,omitempty"`
				Unit                string `json:"unit" bson:"unit,omitempty"`
				SoilType            string `json:"soilType" bson:"soilType,omitempty"`
				LandOwnership       string `json:"landOwnership" bson:"landOwnership,omitempty"`
				CultivationPractice string `json:"cultivationPractice" bson:"cultivationPractice,omitempty"`
				LandPosition        string `json:"landPosition" bson:"landPosition,omitempty"`
				LandType            string `json:"landType" bson:"landType,omitempty"`
				Save                string `json:"save" bson:"save,omitempty"`
				Close               string `json:"close" bson:"close,omitempty"`
			} `json:"addLand" bson:"addLand,omitempty"`
			Livestock struct {
				LiveStock string `json:"liveStock" bson:"liveStock,omitempty"`
				Variety   string `json:"variety" bson:"variety,omitempty"`
				Stage     string `json:"stage" bson:"stage,omitempty"`
				Quanitity string `json:"quanitity" bson:"quanitity,omitempty"`
				Status    string `json:"status" bson:"status,omitempty"`
				Controls  string `json:"controls" bson:"controls,omitempty"`
			} `json:"livestock" bson:"livestock,omitempty"`
			Crop struct {
				Crop          string `json:"crop" bson:"crop,omitempty"`
				InterCrop     string `json:"interCrop" bson:"interCrop,omitempty"`
				Variety       string `json:"variety" bson:"variety,omitempty"`
				Season        string `json:"season" bson:"season,omitempty"`
				Area          string `json:"area" bson:"area,omitempty"`
				Irrigation    string `json:"irrigation" bson:"irrigation,omitempty"`
				StartDate     string `json:"startDate" bson:"startDate,omitempty"`
				CompletedDate string `json:"completedDate" bson:"completedDate,omitempty"`
				Yield         string `json:"yield" bson:"yield,omitempty"`
				Controls      string `json:"controls" bson:"controls,omitempty"`
			} `json:"crop" bson:"crop,omitempty"`
		} `json:"farmerList" bson:"farmerList,omitempty"`
		FeedbackQueries struct {
			FeedbackQueries string `json:"feedbackQueries" bson:"feedbackQueries,omitempty"`
		} `json:"feedbackQueries" bson:"feedbackQueries,omitempty"`
		FarmerAgguregation struct {
			FarmerAgguregation              string `json:"farmerAgguregation" bson:"farmerAgguregation,omitempty"`
			Notes                           string `json:"notes" bson:"notes,omitempty"`
			DownloadDemoContent             string `json:"downloadDemoContent" bson:"downloadDemoContent,omitempty"`
			FarmerUploadDemo                string `json:"farmerUploadDemo" bson:"farmerUploadDemo,omitempty"`
			LocationReference               string `json:"locationReference" bson:"locationReference,omitempty"`
			OrganisationAndProjectReference string `json:"organisationAndProjectReference" bson:"organisationAndProjectReference,omitempty"`
			FarmerUpload                    string `json:"farmerUpload" bson:"farmerUpload,omitempty"`
			Upload                          string `json:"upload" bson:"upload,omitempty"`
			Reset                           string `json:"reset" bson:"reset,omitempty"`
			Browser                         string `json:"browser" bson:"browser,omitempty"`
			SNo                             string `json:"sNo" bson:"sNo,omitempty"`
			MobileNumber                    string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Error                           string `json:"error" bson:"error,omitempty"`
		} `json:"farmerAgguregation" bson:"farmerAgguregation,omitempty"`
		MicroNutrients struct {
			MicroNutrients string `json:"microNutrients" bson:"microNutrients,omitempty"`
			Name           string `json:"name" bson:"name,omitempty"`
			Code           string `json:"code" bson:"code,omitempty"`
			MinValue       string `json:"minValue" bson:"minValue,omitempty"`
			MaxValue       string `json:"maxValue" bson:"maxValue,omitempty"`
			Status         string `json:"status" bson:"status,omitempty"`
			Controls       string `json:"controls" bson:"controls,omitempty"`
		} `json:"microNutrients" bson:"microNutrients,omitempty"`
		MacroNutrients struct {
			MacroNutrients string `json:"macroNutrients" bson:"macroNutrients,omitempty"`
			Name           string `json:"name" bson:"name,omitempty"`
			Code           string `json:"code" bson:"code,omitempty"`
			MinValue       string `json:"minValue" bson:"minValue,omitempty"`
			MaxValue       string `json:"maxValue" bson:"maxValue,omitempty"`
			Status         string `json:"status" bson:"status,omitempty"`
			Controls       string `json:"controls" bson:"controls,omitempty"`
		} `json:"macroNutrients" bson:"macroNutrients,omitempty"`
		LandUpload struct {
			LandUpload       string `json:"landUpload" bson:"landUpload,omitempty"`
			FarmerLandUpload string `json:"farmerLandUpload" bson:"farmerLandUpload,omitempty"`
			Index            string `json:"index" bson:"index,omitempty"`
			Download         string `json:"download" bson:"download,omitempty"`
			Reset            string `json:"reset" bson:"reset,omitempty"`
			Result           string `json:"result" bson:"result,omitempty"`
			Totalrows        string `json:"totalrows" bson:"totalrows,omitempty"`
			Success          string `json:"success" bson:"success,omitempty"`
			Failure          string `json:"failure" bson:"failure,omitempty"`
			Errorreport      string `json:"errorreport" bson:"errorreport,omitempty"`
			Rowno            string `json:"rowno" bson:"rowno,omitempty"`
			Reason           string `json:"reason" bson:"reason,omitempty"`
			Browse           string `json:"browse" bson:"browse,omitempty"`
		} `json:"landUpload" bson:"landUpload,omitempty"`
		CropUpload struct {
			CropUpload       string `json:"cropUpload" bson:"cropUpload,omitempty"`
			FarmerCropUpload string `json:"farmerCropUpload" bson:"farmerCropUpload,omitempty"`
		} `json:"cropUpload" bson:"cropUpload,omitempty"`
		FarmerSoilUpload struct {
			FarmerSoilUpload string `json:"farmerSoilUpload" bson:"farmerSoilUpload,omitempty"`
		} `json:"farmerSoilUpload" bson:"farmerSoilUpload,omitempty"`
		Caste struct {
			Caste    string `json:"caste" bson:"caste,omitempty"`
			AddCaste string `json:"addCaste" bson:"addCaste,omitempty"`
			Name     string `json:"name" bson:"name,omitempty"`
			Status   string `json:"status" bson:"status,omitempty"`
			Controls string `json:"controls" bson:"controls,omitempty"`
		} `json:"caste" bson:"caste,omitempty"`
		FarmerCasteUpdate struct {
			FarmerCasteUpdate string `json:"farmerCasteUpdate" bson:"farmerCasteUpdate,omitempty"`
		} `json:"farmerCasteUpdate" bson:"farmerCasteUpdate,omitempty"`
		LocateFarmer struct {
			LocateFarmer string `json:"locateFarmer" bson:"locateFarmer,omitempty"`
			State        string `json:"state" bson:"state,omitempty"`
			District     string `json:"district" bson:"district,omitempty"`
			Block        string `json:"block" bson:"block,omitempty"`
			Grampanchyat string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
			Village      string `json:"village" bson:"village,omitempty"`
			Reset        string `json:"reset" bson:"reset,omitempty"`
			Search       string `json:"search" bson:"search,omitempty"`
			Track        string `json:"track" bson:"track,omitempty"`
		} `json:"locateFarmer" bson:"locateFarmer,omitempty"`
	} `json:"farmers" bson:"farmers,omitempty"`
	KnowledgeDomain struct {
		KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
		ContentCreation struct {
			ContentCreation string `json:"contentCreation" bson:"contentCreation,omitempty"`
			Sms             string `json:"sms" bson:"sms,omitempty"`
			Voice           string `json:"voice" bson:"voice,omitempty"`
			Video           string `json:"video" bson:"video,omitempty"`
			Poster          string `json:"poster" bson:"poster,omitempty"`
			Document        string `json:"document" bson:"document,omitempty"`
			Compedium       string `json:"compedium" bson:"compedium,omitempty"`
			Create          struct {
				Organisation                  string `json:"organisation" bson:"organisation,omitempty"`
				Project                       string `json:"project" bson:"project,omitempty"`
				Title                         string `json:"title" bson:"title,omitempty"`
				SubDomain                     string `json:"subDomain" bson:"subDomain,omitempty"`
				Topic                         string `json:"topic" bson:"topic,omitempty"`
				SubTopic                      string `json:"subTopic" bson:"subTopic,omitempty"`
				State                         string `json:"state" bson:"state,omitempty"`
				District                      string `json:"district" bson:"district,omitempty"`
				Block                         string `json:"block" bson:"block,omitempty"`
				Grampanchyat                  string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village                       string `json:"village" bson:"village,omitempty"`
				IgnoreState                   string `json:"ignorestate" bson:"ignorestate,omitempty"`
				IgnoreDistrict                string `json:"ignoredistrict" bson:"ignoredistrict,omitempty"`
				IgnoreBlock                   string `json:"ignoreblock" bson:"ignoreblock,omitempty"`
				IgnoreGrampanchyat            string `json:"ignoregrampanchyat" bson:"ignoregrampanchyat,omitempty"`
				IgnoreVillage                 string `json:"ignorevillage" bson:"ignorevillage,omitempty"`
				AllState                      string `json:"allstate" bson:"allstate,omitempty"`
				AllDistrict                   string `json:"alldistrict" bson:"alldistrict,omitempty"`
				AllBlock                      string `json:"allblock" bson:"allblock,omitempty"`
				AllGrampanchyat               string `json:"allgrampanchyat" bson:"allgrampanchyat,omitempty"`
				AllVillage                    string `json:"allvillage" bson:"allvillage,omitempty"`
				Sms                           string `json:"sms" bson:"sms,omitempty"`
				CreateOnePage                 string `json:"createOnePage" bson:"createOnePage,omitempty"`
				VoiceText                     string `json:"voiceText" bson:"voiceText,omitempty"`
				Record                        string `json:"record" bson:"record,omitempty"`
				FileUpload                    string `json:"fileUpload" bson:"fileUpload,omitempty"`
				Youtube                       string `json:"youtube" bson:"youtube,omitempty"`
				ExternalLink                  string `json:"externalLink" bson:"externalLink,omitempty"`
				VideoUpload                   string `json:"videoUpload" bson:"videoUpload,omitempty"`
				ChooseFile                    string `json:"chooseFile" bson:"chooseFile,omitempty"`
				Submit                        string `json:"submit" bson:"submit,omitempty"`
				UploadFile                    string `json:"uploadFile" bson:"uploadFile,omitempty"`
				NoFileChosen                  string `json:"noFileChosen" bson:"noFileChosen,omitempty"`
				Date                          string `json:"date" bson:"date,omitempty"`
				Mandatoryfields               string `json:"mandatoryfields" bson:"mandatoryfields,omitempty"`
				PleaseCheckallMandatoryfields string `json:"pleaseCheckallMandatoryfields" bson:"pleaseCheckallMandatoryfields,omitempty"`
				ContentCount                  string `json:"contentCount" bson:"contentCount,omitempty"`
				QuiuckLink                    string `json:"quiuckLink" bson:"quiuckLink,omitempty"`
				WeatherData                   string `json:"weatherData" bson:"weatherData,omitempty"`
				RssFeed                       string `json:"rssFeed" bson:"rssFeed,omitempty"`
				BannedItems                   string `json:"bannedItems" bson:"bannedItems,omitempty"`
				RelatedContent                string `json:"relatedContent" bson:"relatedContent,omitempty"`
				GetDissemination              string `json:"getDissemination" bson:"getDissemination,omitempty"`
				Farmer                        string `json:"farmer" bson:"farmer,omitempty"`
				User                          string `json:"user" bson:"user,omitempty"`
			} `json:"create" bson:"create,omitempty"`
		} `json:"contentCreation" bson:"contentCreation,omitempty"`
		ContentManager struct {
			ContentManager     string `json:"contentManager" bson:"contentManager,omitempty"`
			CreatedSms         string `json:"createdSms" bson:"createdSms,omitempty"`
			ApprovedSms        string `json:"approvedSms" bson:"approvedSms,omitempty"`
			TranslatedSms      string `json:"translatedSms" bson:"translatedSms,omitempty"`
			PublishedSms       string `json:"publishedSms" bson:"publishedSms,omitempty"`
			RejectedSms        string `json:"rejectedSms" bson:"rejectedSms,omitempty"`
			CreatedPoster      string `json:"createdPoster" bson:"createdPoster,omitempty"`
			ApprovedPoster     string `json:"approvedPoster" bson:"approvedPoster,omitempty"`
			TranslatedPoster   string `json:"translatedPoster" bson:"translatedPoster,omitempty"`
			PublishedPoster    string `json:"publishedPoster" bson:"publishedPoster,omitempty"`
			RejectedPoster     string `json:"rejectedPoster" bson:"rejectedPoster,omitempty"`
			CreatedVoice       string `json:"createdVoice" bson:"createdVoice,omitempty"`
			ApprovedVoice      string `json:"approvedVoice" bson:"approvedVoice,omitempty"`
			TranslatedVoice    string `json:"translatedVoice" bson:"translatedVoice,omitempty"`
			PublishedVoice     string `json:"publishedVoice" bson:"publishedVoice,omitempty"`
			RejectedVoice      string `json:"rejectedVoice" bson:"rejectedVoice,omitempty"`
			CreatedVideo       string `json:"createdVideo" bson:"createdVideo,omitempty"`
			ApprovedVideo      string `json:"approvedVideo" bson:"approvedVideo,omitempty"`
			TranslatedVideo    string `json:"translatedVideo" bson:"translatedVideo,omitempty"`
			PublishedVideo     string `json:"publishedVideo" bson:"publishedVideo,omitempty"`
			RejectedVideo      string `json:"rejectedVideo" bson:"rejectedVideo,omitempty"`
			CreatedDocument    string `json:"createdDocument" bson:"createdDocument,omitempty"`
			ApprovedDocument   string `json:"approvedDocument" bson:"approvedDocument,omitempty"`
			TranslatedDocument string `json:"translatedDocument" bson:"translatedDocument,omitempty"`
			PublishedDocument  string `json:"publishedDocument" bson:"publishedDocument,omitempty"`
			RejectedDocument   string `json:"rejectedDocument" bson:"rejectedDocument,omitempty"`
			ReviewedBy         string `json:"reviewedBy" bson:"reviewedBy,omitempty"`
			CreatedBy          string `json:"createdBy" bson:"createdBy,omitempty"`
			View               struct {
				BasicInfromation string `json:"basicInfromation" bson:"basicInfromation,omitempty"`
				Organisation     string `json:"organisation" bson:"organisation,omitempty"`
				Project          string `json:"project" bson:"project,omitempty"`
				KnowledgeDomain  string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
				SubDomain        string `json:"subDomain" bson:"subDomain,omitempty"`
				Topic            string `json:"topic" bson:"topic,omitempty"`
				SubTopic         string `json:"subTopic" bson:"subTopic,omitempty"`
				State            string `json:"state" bson:"state,omitempty"`
				District         string `json:"district" bson:"district,omitempty"`
				Block            string `json:"block" bson:"block,omitempty"`
				Grampanchyat     string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village          string `json:"village" bson:"village,omitempty"`
				Comment          string `json:"comment" bson:"comment,omitempty"`
				Post             string `json:"post" bson:"post,omitempty"`
				SmsContent       string `json:"smsContent" bson:"smsContent,omitempty"`
				PosterContent    string `json:"posterContent" bson:"posterContent,omitempty"`
				VoiceContent     string `json:"voiceContent" bson:"voiceContent,omitempty"`
				VideoContent     string `json:"videoContent" bson:"videoContent,omitempty"`
				DocumentContent  string `json:"documentContent" bson:"documentContent,omitempty"`
				Content          struct {
					Title            string `json:"title" bson:"title,omitempty"`
					OtherTranslation string `json:"otherTranslation" bson:"otherTranslation,omitempty"`
					Translated       string `json:"translated" bson:"translated,omitempty"`
					Published        string `json:"published" bson:"published,omitempty"`
					Rejected         string `json:"rejected" bson:"rejected,omitempty"`
					Deleted          string `json:"deleted" bson:"deleted,omitempty"`
				} `json:"content" bson:"content,omitempty"`
				Feedback struct {
					FeedBack     string `json:"feedBack" bson:"feedBack,omitempty"`
					FeedbackType string `json:"feedbackType" bson:"feedbackType,omitempty"`
					Farmer       string `json:"farmer" bson:"farmer,omitempty"`
					Village      string `json:"village" bson:"village,omitempty"`
					Date         string `json:"date" bson:"date,omitempty"`
					Rating       string `json:"rating" bson:"rating,omitempty"`
				} `json:"feedback" bson:"feedback,omitempty"`
				Translation struct {
					ContentTitle string `json:"contentTitle" bson:"contentTitle,omitempty"`
					Language     string `json:"language" bson:"language,omitempty"`
					Save         string `json:"save" bson:"save,omitempty"`
				} `json:"translation" bson:"translation,omitempty"`
			} `json:"view" bson:"view,omitempty"`
		} `json:"contentManager" bson:"contentManager,omitempty"`
		TranslateReview struct {
			TranslateReview string `json:"translateReview" bson:"translateReview,omitempty"`
			AllContent      string `json:"allContent" bson:"allContent,omitempty"`
			Sms             string `json:"sms" bson:"sms,omitempty"`
			Voice           string `json:"voice" bson:"voice,omitempty"`
			Poster          string `json:"poster" bson:"poster,omitempty"`
			Translated      string `json:"translated" bson:"translated,omitempty"`
			Published       string `json:"published" bson:"published,omitempty"`
			Rejected        string `json:"rejected" bson:"rejected,omitempty"`
			Deleted         string `json:"deleted" bson:"deleted,omitempty"`
			ContentId       string `json:"contentId" bson:"contentId,omitempty"`
			Type            string `json:"type" bson:"type,omitempty"`
			Content         string `json:"content" bson:"content,omitempty"`
			Org             string `json:"org" bson:"org,omitempty"`
			Lang            string `json:"lang" bson:"lang,omitempty"`
			TranslateBy     string `json:"translateBy" bson:"translateBy,omitempty"`
			TranslateDate   string `json:"translateDate" bson:"translateDate,omitempty"`
			ReviewedBy      string `json:"reviewedBy" bson:"reviewedBy,omitempty"`
			ReviewedDate    string `json:"reviewedDate" bson:"reviewedDate,omitempty"`
			Action          string `json:"action" bson:"action,omitempty"`
		} `json:"translateReview" bson:"translateReview,omitempty"`
		ContentAgguregation struct {
			ContentAgguregation string `json:"contentAgguregation" bson:"contentAgguregation,omitempty"`
			ContentUpload       string `json:"contentUpload" bson:"contentUpload,omitempty"`
			Upload              string `json:"upload" bson:"upload,omitempty"`
			Reset               string `json:"reset" bson:"reset,omitempty"`
			Browse              string `json:"browse" bson:"browse,omitempty"`
			Note                string `json:"note" bson:"note,omitempty"`
		} `json:"contentAgguregation" bson:"contentAgguregation,omitempty"`
		ContentSearch struct {
			ContentSearch string `json:"contentSearch" bson:"contentSearch,omitempty"`
			Table         struct {
				ContentId   string `json:"contentId" bson:"contentId,omitempty"`
				Content     string `json:"content" bson:"content,omitempty"`
				Type        string `json:"type" bson:"type,omitempty"`
				District    string `json:"district" bson:"district,omitempty"`
				Date        string `json:"date" bson:"date,omitempty"`
				Source      string `json:"source" bson:"source,omitempty"`
				Status      string `json:"status" bson:"status,omitempty"`
				UserVisit   string `json:"userVisit" bson:"userVisit,omitempty"`
				FarmerVisit string `json:"farmerVisit" bson:"farmerVisit,omitempty"`
				GuestVisit  string `json:"guestVisit" bson:"guestVisit,omitempty"`
				Control     string `json:"control" bson:"control,omitempty"`
			} `json:"table" bson:"table,omitempty"`
			Filter struct {
				SelectOption    string `json:"selectOption" bson:"selectOption,omitempty"`
				Type            string `json:"type" bson:"type,omitempty"`
				Source          string `json:"source" bson:"source,omitempty"`
				Status          string `json:"status" bson:"status,omitempty"`
				State           string `json:"state" bson:"state,omitempty"`
				Organisation    string `json:"organisation" bson:"organisation,omitempty"`
				KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
				SoilType        string `json:"soilType" bson:"soilType,omitempty"`
				ContentType     string `json:"contentType" bson:"contentType,omitempty"`
				ContentTitle    string `json:"contentTitle" bson:"contentTitle,omitempty"`
				Comment         string `json:"comment" bson:"comment,omitempty"`
				Classification  string `json:"classification" bson:"classification,omitempty"`
				Category        string `json:"category" bson:"category,omitempty"`
				SubCategory     string `json:"subCategory" bson:"subCategory,omitempty"`
				Commodity       string `json:"commodity" bson:"commodity,omitempty"`
				Cause           string `json:"cause" bson:"cause,omitempty"`
				Controls        string `json:"controls" bson:"controls,omitempty"`
				Reset           string `json:"reset" bson:"reset,omitempty"`
				Search          string `json:"search" bson:"search,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"contentSearch" bson:"contentSearch,omitempty"`
		Dissemination struct {
			Dissemination string `json:"dissemination" bson:"dissemination,omitempty"`
			Organisation  string `json:"organisation" bson:"organisation,omitempty"`
			Project       string `json:"project" bson:"project,omitempty"`
			State         string `json:"state" bson:"state,omitempty"`
			District      string `json:"district" bson:"district,omitempty"`
			Block         string `json:"block" bson:"block,omitempty"`
			From          string `json:"from" bson:"from,omitempty"`
			To            string `json:"to" bson:"to,omitempty"`
			Search        string `json:"search" bson:"search,omitempty"`
			Reset         string `json:"reset" bson:"reset,omitempty"`
			Sms           string `json:"sms" bson:"sms,omitempty"`
			Email         string `json:"email" bson:"email,omitempty"`
			Notification  string `json:"notification" bson:"notification,omitempty"`
			Whatsapp      string `json:"whatsapp" bson:"whatsapp,omitempty"`
			Telegram      string `json:"telegram" bson:"telegram,omitempty"`
			list          struct {
				Message           string `json:"message" bson:"message,omitempty"`
				RecordId          string `json:"recordId" bson:"recordId,omitempty"`
				CreatedBy         string `json:"createdBy" bson:"createdBy,omitempty"`
				DisseminationDate string `json:"disseminationDate" bson:"disseminationDate,omitempty"`
				FarmerCount       string `json:"farmerCount" bson:"farmerCount,omitempty"`
				UserCount         string `json:"userCount" bson:"userCount,omitempty"`
			}
		} `json:"dissemination" bson:"dissemination,omitempty"`
		Gis struct {
			Gis         string `json:"gis" bson:"gis,omitempty"`
			State       string `json:"state" bson:"state,omitempty"`
			District    string `json:"district" bson:"district,omitempty"`
			AidCategory string `json:"aidCategory" bson:"aidCategory,omitempty"`
			Reset       string `json:"reset" bson:"reset,omitempty"`
			Search      string `json:"search" bson:"search,omitempty"`
		} `json:"gis" bson:"gis,omitempty"`
		Embed struct {
			Embed string `json:"embed" bson:"embed,omitempty"`
		} `json:"embed" bson:"embed,omitempty"`
	} `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
	ImportantLinks struct {
		ImportantLinks string `json:"importantLinks" bson:"importantLinks,omitempty"`
		CreateSms      string `json:"createSms" bson:"createSms,omitempty"`
		CreateVoice    string `json:"createVoice" bson:"createVoice,omitempty"`
		CreateVideo    string `json:"createvideo" bson:"createvideo,omitempty"`
		CreatePoster   string `json:"createPoster" bson:"createPoster,omitempty"`
		CreateDocument string `json:"createDocument" bson:"createDocument,omitempty"`
		ContentSearch  string `json:"contentSearch" bson:"contentSearch,omitempty"`
	} `json:"importantLinks" bson:"importantLinks,omitempty"`
	FeedBacksAndQueries struct {
		FeedBacksAndQueries string `json:"feedBacksAndQueries" bson:"feedBacksAndQueries,omitempty"`
		FeedBack            struct {
			FeedBack    string `json:"feedBack" bson:"feedBack,omitempty"`
			AddFeedback struct {
				FeedbackType    string `json:"feedbackType" bson:"feedbackType,omitempty"`
				Farmer          string `json:"farmer" bson:"farmer,omitempty"`
				Anonymous       string `json:"anonymous" bson:"anonymous,omitempty"`
				State           string `json:"state" bson:"state,omitempty"`
				District        string `json:"district" bson:"district,omitempty"`
				Block           string `json:"block" bson:"block,omitempty"`
				Grampanchyat    string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village         string `json:"village" bson:"village,omitempty"`
				KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
				SubDomain       string `json:"subDomain" bson:"subDomain,omitempty"`
				Title           string `json:"title" bson:"title,omitempty"`
				Relevance       string `json:"relevance" bson:"relevance,omitempty"`
				Timeliness      string `json:"timeliness" bson:"timeliness,omitempty"`
				Completeness    string `json:"completeness" bson:"completeness,omitempty"`
				Understandable  string `json:"understandable" bson:"understandable,omitempty"`
			} `json:"addFeedback" bson:"addFeedback,omitempty"`
			Filter struct {
				State        string `json:"state" bson:"state,omitempty"`
				District     string `json:"district" bson:"district,omitempty"`
				Block        string `json:"block" bson:"block,omitempty"`
				Grampanchyat string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village      string `json:"village" bson:"village,omitempty"`
				FromDate     string `json:"fromDate" bson:"fromDate,omitempty"`
				ToDate       string `json:"toDate" bson:"toDate,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
			List struct {
				FeedBackType  string `json:"feedBackType" bson:"feedBackType,omitempty"`
				RecordId      string `json:"recordId" bson:"recordId,omitempty"`
				Query         string `json:"query" bson:"query,omitempty"`
				Farmer        string `json:"farmer" bson:"farmer,omitempty"`
				Village       string `json:"village" bson:"village,omitempty"`
				Feedback      string `json:"feedback" bson:"feedback,omitempty"`
				Date          string `json:"date" bson:"date,omitempty"`
				OverallRating string `json:"overallRating" bson:"overallRating,omitempty"`
				Type          string `json:"type" bson:"type,omitempty"`
				Status        string `json:"status" bson:"status,omitempty"`
			} `json:"list" bson:"list,omitempty"`
		} `json:"feedBack" bson:"feedBack,omitempty"`
		Queries struct {
			Queries      string `json:"queries" bson:"queries,omitempty"`
			UnResolved   string `json:"unResolved" bson:"unResolved,omitempty"`
			AssignToMe   string `json:"assignToMe" bson:"assignToMe,omitempty"`
			Assigned     string `json:"assigned" bson:"assigned,omitempty"`
			ResolvedByMe string `json:"resolvedByMe" bson:"resolvedByMe,omitempty"`
			Resolved     string `json:"resolved" bson:"resolved,omitempty"`
			AddQuery     struct {
				QueryType    string `json:"queryType" bson:"queryType,omitempty"`
				Anonymous    string `json:"anonymous" bson:"anonymous,omitempty"`
				CommonQuery  string `json:"commonQuery" bson:"commonQuery,omitempty"`
				Farmer       string `json:"farmer" bson:"farmer,omitempty"`
				State        string `json:"state" bson:"state,omitempty"`
				District     string `json:"district" bson:"district,omitempty"`
				Block        string `json:"block" bson:"block,omitempty"`
				Grampanchyat string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village      string `json:"village" bson:"village,omitempty"`
				Title        string `json:"title" bson:"title,omitempty"`
				Query        string `json:"query" bson:"query,omitempty"`
				Download     string `json:"download" bson:"download,omitempty"`
				AddFiles     string `json:"addFiles" bson:"addFiles,omitempty"`
				DropFile     string `json:"dropFile" bson:"dropFile,omitempty"`
				Save         string `json:"save" bson:"save,omitempty"`
				Close        string `json:"close" bson:"close,omitempty"`
			} `json:"addQuery" bson:"addQuery,omitempty"`
			List struct {
				Type         string `json:"type" bson:"type,omitempty"`
				Farmer       string `json:"farmer" bson:"farmer,omitempty"`
				Village      string `json:"village" bson:"village,omitempty"`
				KD           string `json:"kd" bson:"kd,omitempty"`
				SD           string `json:"sd" bson:"sd,omitempty"`
				State        string `json:"state" bson:"state,omitempty"`
				District     string `json:"district" bson:"district,omitempty"`
				CreatedBy    string `json:"createdBy" bson:"createdBy,omitempty"`
				ResolvedDate string `json:"resolvedDate" bson:"resolvedDate,omitempty"`
				CreatedDate  string `json:"createdDate" bson:"createdDate,omitempty"`
				AssignedDate string `json:"assignedDate" bson:"assignedDate,omitempty"`
				AssignedTo   string `json:"assignedTo" bson:"assignedTo,omitempty"`
				ResolvedBy   string `json:"resolvedBy" bson:"resolvedBy,omitempty"`
				Query        string `json:"query" bson:"query,omitempty"`
				Image        string `json:"image" bson:"image,omitempty"`
				Control      string `json:"control" bson:"control,omitempty"`
			} `json:"list" bson:"list,omitempty"`
			Assign struct {
				AssignUser string `json:"assignUser" bson:"assignUser,omitempty"`
				User       string `json:"user" bson:"user,omitempty"`
				AddUser    string `json:"addUser" bson:"addUser,omitempty"`
				Close      string `json:"close" bson:"close,omitempty"`
			} `json:"assign" bson:"assign,omitempty"`
			Edit struct {
				KnowledgeDomain string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
				SubDomain       string `json:"subDomain" bson:"subDomain,omitempty"`
				Title           string `json:"title" bson:"title,omitempty"`
				Query           string `json:"query" bson:"query,omitempty"`
				Download        string `json:"download" bson:"download,omitempty"`
				Update          string `json:"update" bson:"update,omitempty"`
				Close           string `json:"close" bson:"close,omitempty"`
			} `json:"edit" bson:"edit,omitempty"`
			Resolve struct {
				Query        string `json:"query" bson:"query,omitempty"`
				Solution     string `json:"solution" bson:"solution,omitempty"`
				Download     string `json:"download" bson:"download,omitempty"`
				ChooseFile   string `json:"chooseFile" bson:"chooseFile,omitempty"`
				NoFileChosen string `json:"noFileChosen" bson:"noFileChosen,omitempty"`
				Resolve      string `json:"resolve" bson:"resolve,omitempty"`
				Close        string `json:"close" bson:"close,omitempty"`
			} `json:"resolve" bson:"resolve,omitempty"`
			Filter struct {
				State        string `json:"state" bson:"state,omitempty"`
				District     string `json:"district" bson:"district,omitempty"`
				Block        string `json:"block" bson:"block,omitempty"`
				Grampanchyat string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
				Village      string `json:"village" bson:"village,omitempty"`
				Search       string `json:"search" bson:"search,omitempty"`
				Reset        string `json:"reset" bson:"reset,omitempty"`
				ShowFilter   string `json:"showFilter" bson:"showFilter,omitempty"`
				CloseFilter  string `json:"closeFilter" bson:"closeFilter,omitempty"`
				SearchResult string `json:"searchResult" bson:"searchResult,omitempty"`
				ResultsFound string `json:"resultsFound" bson:"resultsFound,omitempty"`
			} `json:"filter" bson:"filter,omitempty"`
		} `json:"queries" bson:"queries,omitempty"`
	} `json:"feedBacksAndQueries" bson:"feedBacksAndQueries,omitempty"`
	Report struct {
		Report     string `json:"report" bson:"report,omitempty"`
		UserReport struct {
			UserReport        string `json:"userReport" bson:"userReport,omitempty"`
			UserRole          string `json:"userRole" bson:"userRole,omitempty"`
			AccessLevel       string `json:"accessLevel" bson:"accessLevel,omitempty"`
			State             string `json:"state" bson:"state,omitempty"`
			District          string `json:"district" bson:"district,omitempty"`
			Search            string `json:"search" bson:"search,omitempty"`
			Reset             string `json:"reset" bson:"reset,omitempty"`
			IgnoreRole        string `json:"ignoreRole" bson:"ignoreRole,omitempty"`
			IgnoreAccessLevel string `json:"ignoreAccessLevel" bson:"ignoreAccessLevel,omitempty"`
			SelectRole        string `json:"selectRole" bson:"selectRole,omitempty"`
			SelectAccessLevel string `json:"selectAccessLevel" bson:"selectAccessLevel,omitempty"`
			Country           string `json:"country" bson:"country,omitempty"`
		} `json:"userReport" bson:"userReport,omitempty"`
		FarmerReport struct {
			FarmerReport string `json:"farmerReport" bson:"farmerReport,omitempty"`
			State        string `json:"state" bson:"state,omitempty"`
			District     string `json:"district" bson:"district,omitempty"`
			Block        string `json:"block" bson:"block,omitempty"`
			Grampanchyat string `json:"grampanchyat" bson:"grampanchyat,omitempty"`
			Village      string `json:"village" bson:"village,omitempty"`
			CreatedFrom  string `json:"createdFrom" bson:"createdFrom,omitempty"`
			CreatedTo    string `json:"createdTo" bson:"createdTo,omitempty"`
			Search       string `json:"search" bson:"search,omitempty"`
			Reset        string `json:"reset" bson:"reset,omitempty"`
		} `json:"farmerReport" bson:"farmerReport,omitempty"`
		DuplicateUserReport struct {
			DuplicateUserReport string `json:"duplicateUserReport" bson:"duplicateUserReport,omitempty"`
			ByMobileNumber      string `json:"byMobileNumber" bson:"byMobileNumber,omitempty"`
			ByEmail             string `json:"byEmail" bson:"byEmail,omitempty"`
			ExportExcel         string `json:"exportExcel" bson:"exportExcel,omitempty"`
			MobileNumber        string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Name                string `json:"name" bson:"name,omitempty"`
			UserName            string `json:"userName" bson:"userName,omitempty"`
			Email               string `json:"email" bson:"email,omitempty"`
			Role                string `json:"role" bson:"role,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			District            string `json:"district" bson:"district,omitempty"`
		} `json:"duplicateUserReport" bson:"duplicateUserReport,omitempty"`
		DuplicateFarmersReport struct {
			DuplicateFarmersReport string `json:"duplicateFarmersReport" bson:"duplicateFarmersReport,omitempty"`
			ByMobileNumber         string `json:"byMobileNumber" bson:"byMobileNumber,omitempty"`
			ByEmail                string `json:"byEmail" bson:"byEmail,omitempty"`
			ExportExcel            string `json:"exportExcel" bson:"exportExcel,omitempty"`
			MobileNumber           string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
			Name                   string `json:"name" bson:"name,omitempty"`
			Email                  string `json:"email" bson:"email,omitempty"`
			State                  string `json:"state" bson:"state,omitempty"`
			District               string `json:"district" bson:"district,omitempty"`
		} `json:"duplicateFarmersReport" bson:"duplicateFarmersReport,omitempty"`
		DuplicateContentReport struct {
			DuplicateContentReport string `json:"duplicateContentReport" bson:"duplicateContentReport,omitempty"`
			SearchResult           string `json:"searchResult" bson:"searchResult,omitempty"`
			ResultsFound           string `json:"resultsFound" bson:"resultsFound,omitempty"`
			KnowledgeDomain        string `json:"knowledgeDomain" bson:"knowledgeDomain,omitempty"`
			SubDomain              string `json:"subDomain" bson:"subDomain,omitempty"`
			Topic                  string `json:"topic" bson:"topic,omitempty"`
			SubTopic               string `json:"subTopic" bson:"subTopic,omitempty"`
			Season                 string `json:"season" bson:"season,omitempty"`
			RecordId               string `json:"recordId" bson:"recordId,omitempty"`
		} `json:"duplicateContentReport" bson:"duplicateContentReport,omitempty"`
		DisseminationReport struct {
			DisseminationReport string `json:"disseminationReport" bson:"disseminationReport,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			From                string `json:"from" bson:"from,omitempty"`
			To                  string `json:"to" bson:"to,omitempty"`
			Search              string `json:"search" bson:"search,omitempty"`
			Reset               string `json:"reset" bson:"reset,omitempty"`
			District            string `json:"district" bson:"district,omitempty"`
			Farmer              string `json:"farmer" bson:"farmer,omitempty"`
			Sms                 string `json:"sms" bson:"sms,omitempty"`
			Voice               string `json:"voice" bson:"voice,omitempty"`
			Video               string `json:"video" bson:"video,omitempty"`
			Poster              string `json:"poster" bson:"poster,omitempty"`
			Document            string `json:"document" bson:"document,omitempty"`
			TotalDisseminated   string `json:"totalDisseminated" bson:"totalDisseminated,omitempty"`
		} `json:"disseminationReport" bson:"disseminationReport,omitempty"`
	} `json:"report" bson:"report,omitempty"`
	Swc struct {
		Swc                       string `json:"swc" bson:"swc,omitempty"`
		Sms                       string `json:"sms" bson:"sms,omitempty"`
		Whatsapp                  string `json:"whatsapp" bson:"whatsapp,omitempty"`
		Email                     string `json:"email" bson:"email,omitempty"`
		Telegram                  string `json:"telegram" bson:"telegram,omitempty"`
		Notification              string `json:"notification" bson:"notification,omitempty"`
		MobileNumber              string `json:"mobileNumber" bson:"mobileNumber,omitempty"`
		Name                      string `json:"name" bson:"name,omitempty"`
		UserName                  string `json:"userName" bson:"userName,omitempty"`
		Usertype                  string `json:"usertype" bson:"usertype,omitempty"`
		Date                      string `json:"date" bson:"date,omitempty"`
		IsJob                     string `json:"isJob" bson:"isJob,omitempty"`
		SendFor                   string `json:"sendFor" bson:"sendFor,omitempty"`
		Message                   string `json:"message" bson:"message,omitempty"`
		CmmunicationCreditManager struct {
			CmmunicationCreditManager string `json:"communicationCreditManager" bson:"communicationCreditManager,omitempty"`
		} `json:"communicationCreditManager" bson:"communicationCreditManager,omitempty"`
	} `json:"swc" bson:"swc,omitempty"`
	AiManagement struct {
		AiManagement     string `json:"aiManagement" bson:"aiManagement,omitempty"`
		MasterManagement struct {
			MasterManagement string `json:"masterManagement" bson:"masterManagement,omitempty"`
			Compendium       struct {
				Compendium              string `json:"compendium" bson:"compendium,omitempty"`
				Addcompendium           string `json:"addcompendium" bson:"addcompendium,omitempty"`
				Question                string `json:"question" bson:"question,omitempty"`
				Answer                  string `json:"answer" bson:"answer,omitempty"`
				Author                  string `json:"author" bson:"author,omitempty"`
				Action                  string `json:"action" bson:"action,omitempty"`
				Close                   string `json:"close" bson:"close,omitempty"`
				Add                     string `json:"add" bson:"add,omitempty"`
				NewCompendiyam          string `json:"newCompendiyam" bson:"newCompendiyam,omitempty"`
				Farmer                  string `json:"farmer" bson:"farmer,omitempty"`
				User                    string `json:"user" bson:"user,omitempty"`
				LastPublishingDate      string `json:"lastPublishingDate" bson:"lastPublishingDate,omitempty"`
				PublishingSendDate      string `json:"publishingSendDate" bson:"publishingSendDate,omitempty"`
				State                   string `json:"state" bson:"state,omitempty"`
				District                string `json:"district" bson:"district,omitempty"`
				Village                 string `json:"village" bson:"village,omitempty"`
				Block                   string `json:"block" bson:"block,omitempty"`
				GramPanchayat           string `json:"gramPanchayat" bson:"gramPanchayat,omitempty"`
				SelectState             string `json:"selectstate" bson:"selectstate,omitempty"`
				SelectDistrict          string `json:"selectdistrict" bson:"selectdistrict,omitempty"`
				SelectVillage           string `json:"selectvillage" bson:"selectvillage,omitempty"`
				SelectBlock             string `json:"selectblock" bson:"selectblock,omitempty"`
				SelectGramPanchayat     string `json:"selectgramPanchayat" bson:"selectgramPanchayat,omitempty"`
				ApprovalType            string `json:"approvalType" bson:"approvalType,omitempty"`
				Automatic               string `json:"automatic" bson:"automatic,omitempty"`
				OnApproval              string `json:"onApproval" bson:"onApproval,omitempty"`
				PublishingOption        string `json:"publishingOption" bson:"publishingOption,omitempty"`
				SendImmediate           string `json:"sendImmediate" bson:"sendImmediate,omitempty"`
				SendOnSpecificDate      string `json:"sendOnSpecificDate" bson:"sendOnSpecificDate,omitempty"`
				Repeat                  string `json:"repeat" bson:"repeat,omitempty"`
				Upload                  string `json:"upload" bson:"upload,omitempty"`
				BulkUpload              string `json:"bulkUpload" bson:"bulkUpload,omitempty"`
				DownloadDemoContentHere string `json:"downloadDemoContentHere" bson:"downloadDemoContentHere,omitempty"`
				English                 string `json:"english" bson:"english,omitempty"`
				Hindi                   string `json:"hindi" bson:"hindi,omitempty"`
				Tamil                   string `json:"tamil" bson:"tamil,omitempty"`
				Marathi                 string `json:"marathi" bson:"marathi,omitempty"`
			} `json:"compendium" bson:"compendium,omitempty"`
			WeatherMaster struct {
				WeatherMaster string `json:"weatherMaster" bson:"weatherMaster,omitempty"`
				State         string `json:"state" bson:"state,omitempty"`
				Action        string `json:"action" bson:"action,omitempty"`
				Minimum       string `json:"minimum" bson:"minimum,omitempty"`
				Medium        string `json:"medium" bson:"medium,omitempty"`
				Disaster      string `json:"disaster" bson:"disaster,omitempty"`
				RainFall      string `json:"rainFall" bson:"rainFall,omitempty"`
				Temperature   string `json:"temperature" bson:"temperature,omitempty"`
				Humidity      string `json:"humidity" bson:"humidity,omitempty"`
				WindSpeed     string `json:"windSpeed" bson:"windSpeed,omitempty"`
				WindDirection string `json:"windDirection" bson:"windDirection,omitempty"`
				Filter        struct {
					SelectSeason string `json:"selectSeason" bson:"selectSeason,omitempty"`
					SelectMonth  string `json:"selectMonth" bson:"selectMonth,omitempty"`
				} `json:"filter" bson:"filter,omitempty"`
			} `json:"weatherMaster" bson:"weatherMaster,omitempty"`
			DistrictWeatherMaster struct {
				DistrictWeatherMaster string `json:"districtWeatherMaster" bson:"districtWeatherMaster,omitempty"`
				District              string `json:"district" bson:"district,omitempty"`
				Action                string `json:"action" bson:"action,omitempty"`
				Minimum               string `json:"minimum" bson:"minimum,omitempty"`
				Medium                string `json:"medium" bson:"medium,omitempty"`
				Disaster              string `json:"disaster" bson:"disaster,omitempty"`
				RainFall              string `json:"rainFall" bson:"rainFall,omitempty"`
				Temperature           string `json:"temperature" bson:"temperature,omitempty"`
				Humidity              string `json:"humidity" bson:"humidity,omitempty"`
				WindSpeed             string `json:"windSpeed" bson:"windSpeed,omitempty"`
				WindDirection         string `json:"windDirection" bson:"windDirection,omitempty"`
				Close                 string `json:"close" bson:"close,omitempty"`
				Savealert             string `json:"savealert" bson:"savealert,omitempty"`
				Addalertdata          string `json:"addalertdata" bson:"addalertdata,omitempty"`
				Chooseyourstate       string `json:"chooseyourstate" bson:"chooseyourstate,omitempty"`
				Selectstate           string `json:"selectstate" bson:"selectstate,omitempty"`
				Allstate              string `json:"allstate" bson:"allstate,omitempty"`
				Choosealerttype       string `json:"choosealerttype" bson:"choosealerttype,omitempty"`
				Selectalerttype       string `json:"selectalerttype" bson:"selectalerttype,omitempty"`
				Choosealertvalue      string `json:"choosealertvalue" bson:"choosealertvalue,omitempty"`
				Selectalertvalue      string `json:"selectalertvalue" bson:"selectalertvalue,omitempty"`
				MinimumValue          string `json:"minimumValue" bson:"minimumValue,omitempty"`
				MediumValue           string `json:"mediumValue" bson:"mediumValue,omitempty"`
				ApprovalType          string `json:"approvalType" bson:"approvalType,omitempty"`
				OnApproval            string `json:"onApproval" bson:"onApproval,omitempty"`
				Add                   string `json:"add" bson:"add,omitempty"`
				Selectyouroption      string `json:"selectyouroption" bson:"selectyouroption,omitempty"`
				Automatic             string `json:"automatic" bson:"automatic,omitempty"`
				Filter                struct {
					SelectSeason string `json:"selectSeason" bson:"selectSeason,omitempty"`
					SelectMonth  string `json:"selectMonth" bson:"selectMonth,omitempty"`
				} `json:"filter" bson:"filter,omitempty"`
			} `json:"districtWeatherMaster" bson:"districtWeatherMaster,omitempty"`
		} `json:"masterManagement" bson:"masterManagement,omitempty"`
		AlertManagement struct {
			AlertManagement string `json:"alertManagement" bson:"alertManagement,omitempty"`
			WeatherAlert    struct {
				WeatherAlert string `json:"weatherAlert" bson:"weatherAlert,omitempty"`
				State        string `json:"state" bson:"state,omitempty"`
				Action       string `json:"action" bson:"action,omitempty"`
				AlertType    string `json:"alertType" bson:"alertType,omitempty"`
				Date         string `json:"date" bson:"date,omitempty"`
				Preview      string `json:"preview" bson:"preview,omitempty"`
				Filter       struct {
					WeatherAlertType string `json:"weatherAlertType" bson:"weatherAlertType,omitempty"`
					WeatherType      string `json:"weatherType" bson:"weatherType,omitempty"`
				} `json:"filter" bson:"filter,omitempty"`
			} `json:"weatherAlert" bson:"weatherAlert,omitempty"`
			DistrictWeatherAlert struct {
				DistrictWeatherAlert string `json:"districtWeatherAlert" bson:"districtWeatherAlert,omitempty"`
				State                string `json:"state" bson:"state,omitempty"`
				District             string `json:"district" bson:"district,omitempty"`
				Action               string `json:"action" bson:"action,omitempty"`
				AlertType            string `json:"alertType" bson:"alertType,omitempty"`
				Date                 string `json:"date" bson:"date,omitempty"`
				Preview              string `json:"preview" bson:"preview,omitempty"`
				Filter               struct {
					WeatherAlertType string `json:"weatherAlertType" bson:"weatherAlertType,omitempty"`
					WeatherType      string `json:"weatherType" bson:"weatherType,omitempty"`
				} `json:"filter" bson:"filter,omitempty"`
			} `json:"districtWeatherAlert" bson:"districtWeatherAlert,omitempty"`
			CompendiumAlert struct {
				CompendiumAlert string `json:"compendiumAlert" bson:"compendiumAlert,omitempty"`
				Problem         string `json:"problem" bson:"problem,omitempty"`
				Solution        string `json:"solution" bson:"solution,omitempty"`
				Action          string `json:"action" bson:"action,omitempty"`
				AlertType       string `json:"alertType" bson:"alertType,omitempty"`
				Farmer          string `json:"farmer" bson:"farmer,omitempty"`
				Preview         string `json:"preview" bson:"preview,omitempty"`
				User            string `json:"user" bson:"user,omitempty"`
				Approved        string `json:"approved" bson:"approved,omitempty"`
				OnePager        string `json:"onePager" bson:"onePager,omitempty"`
			} `json:"compendiumAlert" bson:"compendiumAlert,omitempty"`
			ContigencyAlert struct {
				ContigencyAlert string `json:"contigencyAlert" bson:"contigencyAlert,omitempty"`
				Problem         string `json:"problem" bson:"problem,omitempty"`
				Solution        string `json:"solution" bson:"solution,omitempty"`
				Action          string `json:"action" bson:"action,omitempty"`
				AlertType       string `json:"alertType" bson:"alertType,omitempty"`
				Farmer          string `json:"farmer" bson:"farmer,omitempty"`
				Preview         string `json:"preview" bson:"preview,omitempty"`
				User            string `json:"user" bson:"user,omitempty"`
				Approved        string `json:"approved" bson:"approved,omitempty"`
				OnePager        string `json:"onePager" bson:"onePager,omitempty"`
				Maize           string `json:"maize" bson:"maize,omitempty"`
				PearlMillet     string `json:"pearlMillet" bson:"pearlMillet,omitempty"`
				Onion           string `json:"onion" bson:"onion,omitempty"`
			} `json:"contigencyAlert" bson:"contigencyAlert,omitempty"`
		} `json:"alertManagement" bson:"alertManagement,omitempty"`
		ContigencyMaster struct {
			ContigencyMaster    string `json:"contigencyMaster" bson:"contigencyMaster,omitempty"`
			AddcontigencyMaster string `json:"addcontigencyMaster" bson:"addcontigencyMaster,omitempty"`
			Condion             string `json:"condion" bson:"condion,omitempty"`
			Crop                string `json:"crop" bson:"crop,omitempty"`
			CropStage           string `json:"cropStage" bson:"cropStage,omitempty"`
			ChangeinCrop        string `json:"changeinCrop" bson:"changeinCrop,omitempty"`
			Measure             string `json:"measure" bson:"measure,omitempty"`
			Status              string `json:"status" bson:"status,omitempty"`
			Control             string `json:"controls" bson:"controls,omitempty"`
			State               string `json:"state" bson:"state,omitempty"`
			District            string `json:"district" bson:"district,omitempty"`
			Block               string `json:"block" bson:"block,omitempty"`
			Close               string `json:"close" bson:"close,omitempty"`
			Add                 string `json:"add" bson:"add,omitempty"`
			AddNew              string `json:"addNew" bson:"addNew,omitempty"`
		} `json:"contigencyMaster" bson:"contigencyMaster,omitempty"`
	} `json:"aiManagement" bson:"aiManagement,omitempty"`
	Acl struct {
		Acl     string `json:"acl" bson:"acl,omitempty"`
		Modules struct {
			Modules string `json:"modules" bson:"modules,omitempty"`
		} `json:"modules" bson:"modules,omitempty"`
		Menus struct {
			Menus string `json:"menus" bson:"menus,omitempty"`
		} `json:"menus" bson:"menus,omitempty"`
		Tabs struct {
			Tabs string `json:"tabs" bson:"tabs,omitempty"`
		} `json:"tabs" bson:"tabs,omitempty"`
		Features struct {
			Features string `json:"features" bson:"features,omitempty"`
		} `json:"features" bson:"features,omitempty"`
	} `json:"acl" bson:"acl,omitempty"`
	Uas struct {
		Uas       string `json:"uas" bson:"uas,omitempty"`
		Sms       string `json:"sms" bson:"sms,omitempty"`
		Voice     string `json:"voice" bson:"voice,omitempty"`
		Video     string `json:"video" bson:"video,omitempty"`
		Poster    string `json:"poster" bson:"poster,omitempty"`
		Document  string `json:"document" bson:"document,omitempty"`
		ContentId string `json:"contentId" bson:"contentId,omitempty"`
		Content   string `json:"content" bson:"content,omitempty"`
		Source    string `json:"source" bson:"source,omitempty"`
		Status    string `json:"status" bson:"status,omitempty"`
		ViewMore  string `json:"viewMore" bson:"viewMore,omitempty"`
	} `json:"uas" bson:"uas,omitempty"`
	Pagination struct {
		First        string `json:"first" bson:"first,omitempty"`
		Last         string `json:"last" bson:"last,omitempty"`
		Next         string `json:"next" bson:"next,omitempty"`
		Previous     string `json:"previous" bson:"previous,omitempty"`
		TotalRecords string `json:"totalRecords" bson:"totalRecords,omitempty"`
		TotalPages   string `json:"totalPages" bson:"totalPages,omitempty"`
		Active       string `json:"active" bson:"active,omitempty"`
		Disabled     string `json:"disabled" bson:"disabled,omitempty"`
		NoData       string `json:"noData" bson:"noData,omitempty"`
	} `json:"pagination" bson:"pagination,omitempty"`
	Status  string   `json:"status" form:"status" bson:"status,omitempty"`
	Created *Created `json:"created" form:"created" bson:"created,omitempty"`
}
type CommonLanguageTranslationsFilter struct {
	ActiveStatus []bool   `json:"activestatus,omitempty" form:"activestatus" bson:"activestatus,omitempty"`
	Status       []string `json:"status" form:"status" bson:"status,omitempty"`
	Version      []string `json:"version,omitempty" form:"version" bson:"version,omitempty"`
	SortBy       string   `json:"sortBy"`
	SortOrder    int      `json:"sortOrder"`
	Regex        struct {
		Name string `json:"name" bson:"name"`
	} `json:"regex" bson:"regex"`
}
type RefCommonLanguageTranslations struct {
	CommonLanguageTranslationss `bson:",inline"`
	Ref                         struct {
	} `json:"ref,omitempty" bson:"ref,omitempty"`
}
