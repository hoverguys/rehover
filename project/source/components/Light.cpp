#include "Light.h"

namespace Components {

void Light::Bind(unsigned short slot) {
	GX_LoadLightObj(&lightobj, slot);
}

void PointLight::Setup(Mtx& view, const Transform& transform) {
	guVector pos = transform.position;

	guVecMultiply(view, &pos, &pos);

	GX_InitLightColor(&lightobj, color);
	GX_InitLightPos(&lightobj, pos.x, pos.y, pos.z);
}

void DirectionalLight::Setup(Mtx& view, const Transform& transform) {
	guVector dir = transform.forward;

	guVecMultiplySR(view, &dir, &dir);

	GX_InitLightColor(&lightobj, color);
	GX_InitSpecularDirv(&lightobj, &dir);
	GX_InitLightShininess(&lightobj, shininess);
}

} // namespace Components