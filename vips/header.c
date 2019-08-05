#include "header.h"

unsigned long has_icc_profile(VipsImage *in) {
    return vips_image_get_typeof(in, VIPS_META_ICC_NAME);
}

gboolean remove_icc_profile(VipsImage *in) {
    return vips_image_remove(in, VIPS_META_ICC_NAME);
}

// won't remove the ICC profile
void remove_metadata(VipsImage *in) {
    gchar ** fields = vips_image_get_fields(in);

    for (int i=0; fields[i] != NULL; i++) {
        if (strncmp(fields[i], VIPS_META_ICC_NAME, 16)) {
            vips_image_remove(in, fields[i]);
        }
    }

    g_strfreev(fields);
}

int get_meta_orientation(VipsImage *in) {
	int orientation = 0;
	if (vips_image_get_typeof(in, VIPS_META_ORIENTATION) != 0) {
		vips_image_get_int(in, VIPS_META_ORIENTATION, &orientation);
	}

    return orientation;
}