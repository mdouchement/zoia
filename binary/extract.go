package binary

// ExtractHeader returns the header in Raw format.
func ExtractHeader(patch []byte) Raw {
	return Raw(patch[:SizeHeader])
}

// ExtractModule returns the module in Raw format at the given offset.
func ExtractModule(patch []byte, offset int) Raw {
	l := int(patch[offset])
	return Raw(patch[offset : offset+l*4])
}

// ExtractPatchCable returns the patch cable in Raw format at the given offset.
func ExtractPatchCable(patch []byte, offset int) Raw {
	return Raw(patch[offset : offset+SizePatchCable])
}

// ExtractPage returns the page in Raw format at the given offset.
func ExtractPage(patch []byte, offset int) Raw {
	return Raw(patch[offset : offset+SizePage])
}

// ExtractModuleExtra returns the module extra data in Raw format at the given offset.
// Extra details are stored at the end of the patch in the same order of the modules.
func ExtractModuleExtra(patch []byte, offset int) Raw {
	return Raw(patch[offset : offset+SizeModuleExtra])
}
