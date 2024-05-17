import { z } from "zod";

const FileSchema = z.object({
  name: z.string(),
  path: z.string(),
  size: z.string(),
  extension: z.string().nullable(),
  modifiedTime: z.string(),
});
export const DiskSchema = z.object({
  path: z.string(),
  totalSize: z.string(),
  availableSize: z.string(),
  contents: z.array(FileSchema),
});

export type Disk = z.infer<typeof DiskSchema>;
