#include "Light.h"

namespace Components {

void Light::Bind(unsigned short slot) {
	GX_LoadLightObj(&lightobj, slot);
}

void PointLight::Setup(const Matrix& view, const Transform& transform) {
	Vector pos = view.Multiply(transform.position);

	GX_InitLightColor(&lightobj, color);
	GX_InitLightPosv(&lightobj, &pos);
}

void DirectionalLight::Setup(const Matrix& view, const Transform& transform) {
	Vector pos = view.Multiply(transform.position);
	Vector dir = view.MultiplySR(transform.forward);

	GX_InitLightAttn(&lightobj, 1, 0, 0, 1, 0.1, 0);
	GX_InitLightColor(&lightobj, color);
	GX_InitLightPosv(&lightobj, &pos);
	GX_InitLightDirv(&lightobj, &dir);
	GX_InitLightSpot(&lightobj, 90.0, GX_SP_OFF);
	GX_InitLightDistAttn(&lightobj, 1, 1, GX_DA_OFF);
	GX_InitLightShininess(&lightobj, shininess);
}

} // namespace Components