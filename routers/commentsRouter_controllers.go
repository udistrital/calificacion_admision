package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:CuposPorDependenciaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:DetalleEvaluacionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EntrevistadorEntrevistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EstadoEntrevistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:EvaluacionInscripcionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:RequisitoProgramaAcademicoController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/evaluacion_inscripcion_crud/controllers:TipoEntrevistaController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
