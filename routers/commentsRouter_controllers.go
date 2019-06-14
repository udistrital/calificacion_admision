package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioAdmisionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:CriterioProgramaAcademicoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EntrevistadorEntrevistaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:EstadoEntrevistaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/planesticud/calificacion_admision_crud/controllers:TipoEntrevistaController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
