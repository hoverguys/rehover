#include "Rect.h"

void Rect::Move(const Point2D& delta) {
	start += delta;
	end += delta;
}

void Rect::Resize(const Point2D& size) {
	end += size;
}