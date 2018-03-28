#pragma once

/*!
 * \brief Resource loader
 * Loads a resource from memory
 */
class Resource {
public:
	/*!
	 * \brief Loads a resource from a memory pointer and size
	 *
	 * \param address Pointer to resource data
	 * \param size Size of resource data
	 */
	Resource(void* address, unsigned int size) : address(address), size(size){};
	~Resource() {
		free(address);
	}
	virtual void Initialize() = 0;
	virtual long int ReferenceCount() = 0;

protected:
	friend class ResourceLoader;
	void* address;
	unsigned int size;
};