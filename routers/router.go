// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/descuento_academico_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/requisito",
			beego.NSInclude(
				&controllers.RequisitoController{},
			),
		),

		beego.NSNamespace("/requisito_tipo_descuento",
			beego.NSInclude(
				&controllers.RequisitoTipoDescuentoController{},
			),
		),

		beego.NSNamespace("/tipo_duracion",
			beego.NSInclude(
				&controllers.TipoDuracionController{},
			),
		),

		beego.NSNamespace("/validacion_descuento",
			beego.NSInclude(
				&controllers.ValidacionDescuentoController{},
			),
		),

		beego.NSNamespace("/soporte_descuento",
			beego.NSInclude(
				&controllers.SoporteDescuentoController{},
			),
		),

		beego.NSNamespace("/tipo_descuento",
			beego.NSInclude(
				&controllers.TipoDescuentoController{},
			),
		),

		beego.NSNamespace("/descuentos_dependencia",
			beego.NSInclude(
				&controllers.DescuentosDependenciaController{},
			),
		),

		beego.NSNamespace("/solicitud_descuento",
			beego.NSInclude(
				&controllers.SolicitudDescuentoController{},
			),
		),

		beego.NSNamespace("/tipo_dependencia_descuento",
			beego.NSInclude(
				&controllers.TipoDependenciaDescuentoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
