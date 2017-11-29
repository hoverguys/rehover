#pragma once

template <class T>
class BehaviourSystem : public ex::System<BehaviourSystem<T>> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override {
		es.each<T>([&](ex::Entity entity, T& behaviour) {
			behaviour.Tick(entity, dt);
		});
	}
};