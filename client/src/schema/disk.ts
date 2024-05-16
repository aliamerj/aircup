import { z } from "zod";

export const DiskSchema = z.object({
  disk_path: z.string(),
});

export const DisksSchema = z.array(DiskSchema);

export type Disk = z.infer<typeof DiskSchema>;
export type Disks = z.infer<typeof DisksSchema>;
