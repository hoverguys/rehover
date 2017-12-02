#include "Rect.h"

void Rect::Move(Point2D delta) {
	start += delta;
	end += delta;
}

void Rect::Resize(Point2D size) {
	end += size;
}