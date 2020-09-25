package models

type TrArchivoIcfes struct {
	EvaluacionesInscripcion *[]EvaluacionInscripcion
}

// AddTransaccionArchivoIcfes Transacción para registrar toda la información del archivo del icfes
// func AddTransaccionArchivoIcfes(m *TrArchivoIcfes) (err error) {
// 	o := orm.NewOrm()
// 	err = o.Begin()

// 	for _, v := range *m.EvaluacionesInscripcion {

// 		var evaluacionInscripcion EvaluacionInscripcion
// 		// if errTr := o.QueryTable(new(EvaluacionInscripcion)).RelatedSel().Filter("InscripcionId",v.InscripcionId).Filter("RequisitoProgramaAcademicoId",v.RequisitoProgramaAcademicoId).One(&evaluacionInscripcion); errTr == nil{
// 		if errTr := o.QueryTable(new(EvaluacionInscripcion)).RelatedSel().Filter("InscripcionId", v.InscripcionId).Filter("RequisitoProgramaAcademicoId", v.RequisitoProgramaAcademicoId).One(&evaluacionInscripcion); errTr == nil {
// 			// Si existe hace update
// 			evaluacionInscripcion.NotaFinal = v.NotaFinal
// 			if _, errTr = o.Update(&evaluacionInscripcion, "NotaFinal"); errTr != nil {
// 				err = errTr
// 				fmt.Println(err)
// 				_ = o.Rollback()
// 				return
// 			} else {
// 				// fmt.Println("update para inscripcion",evaluacionInscripcion.InscripcionId,", requisito ",evaluacionInscripcion.RequisitoProgramaAcademicoId,"con valor",evaluacionInscripcion.NotaFinal,"y id",evaluacionInscripcion.Id)
// 			}
// 		} else {
// 			if errTr == orm.ErrNoRows {
// 				if idEvaluacion, errTr := o.Insert(&v); errTr != nil {
// 					err = errTr
// 					fmt.Println(err)
// 					_ = o.Rollback()
// 					return
// 				} else {
// 					// fmt.Println("insert para inscripcion",v.InscripcionId,", requisito ",v.RequisitoProgramaAcademicoId,"con valor",v.NotaFinal,"y id",idEvaluacion)
// 					fmt.Println("insert para inscripcion", v.InscripcionId, ", requisito ", "con valor", v.NotaFinal, "y id", idEvaluacion)
// 				}
// 			} else {
// 				err = errTr
// 				fmt.Println(err)
// 				_ = o.Rollback()
// 				return
// 			}
// 		}
// 	}
// 	_ = o.Commit()
// 	return
// }
