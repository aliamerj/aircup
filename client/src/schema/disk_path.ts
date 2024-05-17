import { z } from "zod";

export const SavedDiskPathSchema = z.object({
  disk_path: z.string(),
});

export const SavedDiskPathsSchema = z.array(SavedDiskPathSchema);
export type SavedDiskPaths = z.infer<typeof SavedDiskPathsSchema>;
