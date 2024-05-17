import { SavedDiskPathsSchema, SavedDiskPaths } from "@/schema/disk_path";

export async function getSavedDisks(): Promise<SavedDiskPaths> {
  try {
    const res = await fetch(import.meta.env.VITE_API_URL + "/disk", {
      method: "GET",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch disks, status: ${res.status}`);
    }

    const data = await res.json();
    const disksPath = SavedDiskPathsSchema.safeParse(data.body);
    if (!disksPath.success) {
      throw new Error("Corrupted data");
    }

    return disksPath.data;
  } catch (error: any) {
    console.error("Error fetching disks:", error);
    return [];
  }
}
