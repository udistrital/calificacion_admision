// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/planesticud/calificacion_admision_crud/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/estado_entrevista",
			beego.NSInclude(
				&controllers.EstadoEntrevistaController{},
			),
		),

		beego.NSNamespace("/tipo_entrevista",
			beego.NSInclude(
				&controllers.TipoEntrevistaController{},
			),
		),

		beego.NSNamespace("/criterio",
			beego.NSInclude(
				&controllers.CriterioController{},
			),
		),

		beego.NSNamespace("/criterio_programa_academico",
			beego.NSInclude(
				&controllers.CriterioProgramaAcademicoController{},
			),
		),

		beego.NSNamespace("/criterio_admision",
			beego.NSInclude(
				&controllers.CriterioAdmisionController{},
			),
		),

		beego.NSNamespace("/entrevistador",
			beego.NSInclude(
				&controllers.EntrevistadorController{},
			),
		),

		beego.NSNamespace("/entrevista",
			beego.NSInclude(
				&controllers.EntrevistaController{},
			),
		),

		beego.NSNamespace("/entrevistador_entrevista",
			beego.NSInclude(
				&controllers.EntrevistadorEntrevistaController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
