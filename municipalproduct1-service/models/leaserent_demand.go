package models


//LeaseRentDemand : ""
type LeaseRentDemand struct {
	RefLeaseRent `bson:",inline"` 
	FY                     []RefLeaseRentDemandFYLog `json:"fy" bson:"fy,omitempty"`
	ProductConfiguration   *RefProductConfiguration   `json:"-" bson:"productConfiguration,omitempty"`	
}


// func (demand *LeaseRentDemand) CalcDemandQuery() ([]bson.M, error) {


	// var mainPipeline []bson.M
	// mainPipeline = append(mainPipeline, bson.M{"$match": bson.M{"uniqueId": mtd.RefLeaseRent.UniqueID}})

	//db.getCollection('leaserent').aggregate([

		//{$match : {uniqueId:"LR00002"}},
// 		mainPipeline = append(mainPipeline,bson.M{
// 			"$lookup":bson.M{
// 			"from":"financialyears",
// 			"let":bson.M{"lrdatefrom":"$dateFrom","lrdateto":"$dateTo","shopcatId":"$shopCategoryId","shopsubcatId":"$shopSubCategoryId"},
// 			"as":"fy",
// 			"pipeline":[]bson.M{
// 			bson.M{"$match":bson.M{"$expr":bson.M{"$and":[]bson.M{
// 				bson.M{"$eq":[]string{"$status","Active"}},
// 				bson.M{"$gte":[]string{"$to","$$lrdatefrom"}},
// 				bson.M{"$eq":[]interface{}{true,bson.M{"$cond":bson.M{"if":bson.M{"$not":[]string{"$$lrdateto"}},"then":true,"else":bson.M{
// 					"$cond":bson.M{"if":bson.M{$gte:[]string{"$$lrdateto","$from"}},"then":true,"else":false}
// 					}}}}},
// 			}}}},

// 			 bson.M{"$sort":bson.M{"order":1}},

// 			  bson.M{
// 				  "$lookup":bson.M{
// 				  "from":"leaserentratemaster",
// 				   "let":bson.M{"tshopcatIds":"$$shopcatId","tshopsubcatIds":"$$shopsubcatId","tfy2":"$to"},
// 				   "as":"ref.rate",
// 			       "pipeline":[]bson.M{
// 				bson.M{"$match":{"$expr":bson.M{"$and":[]bson.M{
// 					 bson.M{"$eq":[]string{"$$tshopcatIds","$shopCategoryId"}},
// 				   bson.M{"$eq":[]string{"$$tshopsubcatIds","$shopSubCategoryId"}},
// 				   bson.M{"$lte":[]string{"doe","$$tfy2"}}
					
// 				   }}}},
// 				bson.M{"$sort":bson.M{"doe":-1}},
					  
// 				  },
				
// 				  },
// 				},
// 				   bson.M{"$addFields":bson.M{"ref.rate":bson.M{"$arrayElemAt":[]interface{}{"$ref.rate",0}}}}
				  
				  
// 		},
// 	},
// }
// 			)
// 			return mainPipeline, nil



// }


// func (demand *LeaseRentDemand) CalcDemand() error {
	
// }