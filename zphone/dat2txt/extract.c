#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define VERSION_LENGTH 4
#define FIRST_INDEX_OFFSET_LENGTH 4
#define INDEX_LENGTH 9

#pragma pack(push, 1)
typedef struct {
    int phone;
    int record_offset;
    unsigned char sim_type;
} Index;
#pragma pack(pop)

int main(int argc, char** argv) {
    FILE* fp = fopen("./phone.dat", "rb");
    if (fp) {
        fseek(fp, 0, SEEK_END);
        long fsize = ftell(fp);
        fseek(fp, 0, SEEK_SET);

        char version[VERSION_LENGTH + 1] = {0};
        fread(version, 1, VERSION_LENGTH, fp);

        int first_index_offset = 0;
        fread(&first_index_offset, FIRST_INDEX_OFFSET_LENGTH, 1, fp);

        long index_count = (fsize - first_index_offset) / INDEX_LENGTH;
        Index* indices = (Index*)malloc(index_count * sizeof(Index));
        if (0 == fseek(fp, first_index_offset, SEEK_SET)) {
			long i;
            for (i = 0; i < index_count; ++i) {
                Index* index = indices + i;
                fread(index, 1, INDEX_LENGTH, fp);
            }

            char buff[1024] = {0};
            for (i = 0; i < index_count; ++i) {
                Index* index = indices + i;
                long offset = index->record_offset;
                if (0 == fseek(fp, offset, SEEK_SET)) {
                    fread(buff, 1, sizeof(buff), fp);
                    printf("%d|%d|%s\n", index->phone, index->sim_type, buff);
                }
            }
        }

        free(indices);
        fclose(fp);
    }

    return 0;
}
