#pragma once
#include "Transform.h"
#include "../pchheader.h"

namespace Components {
class Light {
protected:
	explicit Light(const GXColor color) : color(color){};
	GXLightObj lightobj = {0};

public:
	GXColor color;
	virtual void Setup(const Matrix& view, const Transform& transform) = 0;
	void Bind(unsigned short slot);
};

class PointLight : public Light {
public:
	explicit PointLight(const GXColor color) : Light(color){};
	void Setup(const Matrix& view, const Transform& transform) override;
};

class DirectionalLight : public Light {
public:
	float shininess;
	explicit DirectionalLight(const GXColor color, const float shininess = 0) : Light(color), shininess(shininess){};
	void Setup(const Matrix& view, const Transform& transform) override;
};

} // namespace Components