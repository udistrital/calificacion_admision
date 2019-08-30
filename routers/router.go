// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/evaluacion_inscripcion_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_entrevista",
			beego.NSInclude(
				&controllers.EstadoEntrevistaController{},
			),
		),

		beego.NSNamespace("/requisito",
			beego.NSInclude(
				&controllers.RequisitoController{},
			),
		),

		beego.NSNamespace("/tipo_entrevista",
			beego.NSInclude(
				&controllers.TipoEntrevistaController{},
			),
		),

		beego.NSNamespace("/evaluacion_inscripcion",
			beego.NSInclude(
				&controllers.EvaluacionInscripcionController{},
			),
		),

		beego.NSNamespace("/entrevistador_entrevista",
			beego.NSInclude(
				&controllers.EntrevistadorEntrevistaController{},
			),
		),

		beego.NSNamespace("/entrevistador",
			beego.NSInclude(
				&controllers.EntrevistadorController{},
			),
		),

		beego.NSNamespace("/requisito_programa_academico",
			beego.NSInclude(
				&controllers.RequisitoProgramaAcademicoController{},
			),
		),

		beego.NSNamespace("/cupos_por_dependencia",
			beego.NSInclude(
				&controllers.CuposPorDependenciaController{},
			),
		),

		beego.NSNamespace("/entrevista",
			beego.NSInclude(
				&controllers.EntrevistaController{},
			),
		),

		beego.NSNamespace("/tr_archivo_icfes",
			beego.NSInclude(
				&controllers.TrArchivoIcfesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
