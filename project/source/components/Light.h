#pragma once
#include "Transform.h"
#include <ogc/gx.h>

namespace Components {
class Light {
protected:
	Light(GXColor color) : color(color){};
	GXLightObj lightobj = {0};

public:
	GXColor color;
	virtual void Setup(Matrix view, const Transform& transform) = 0;
	void Bind(unsigned short slot);
};

class PointLight : public Light {
public:
	PointLight(const GXColor color) : Light(color){};
	void Setup(Matrix view, const Transform& transform) override;
};

class DirectionalLight : public Light {
public:
	float shininess;
	DirectionalLight(const GXColor color, const float shininess = 0) : Light(color), shininess(shininess){};
	void Setup(Matrix view, const Transform& transform) override;
};

} // namespace Components