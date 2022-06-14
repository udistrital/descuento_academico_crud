package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:DescuentosDependenciaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:RequisitoTipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SolicitudDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:SoporteDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDependenciaDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:TipoDuracionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/descuento_academico_crud/controllers:ValidacionDescuentoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
