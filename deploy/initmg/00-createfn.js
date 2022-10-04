// db.counters.insert(
//     {
//        _id: "userid",
//        seq: 1
//     },
//     {
//         _id: "incidentid",
//         seq: 1
//     },
//     {
//         _id: "typeid",
//         seq: 7
//     },
//     {
//         _id: "roleid",
//         seq: 3
//     },
//     {
//         _id: "statusid",
//         seq: 2
//     }
//  )
 

// function getNextSequence(name) {
//     var ret = db.counters.findAndModify(
//            {
//              query: { _id: name },
//              update: { $inc: { seq: 1 } },
//              new: true
//            }
//     );
 
//     return ret.seq;
// }

