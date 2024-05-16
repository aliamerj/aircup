import { Disks, DisksSchema } from "@/schema/disk";

export async function getSavedDisks(): Promise<Disks> {
  try {
    const res = await fetch(import.meta.env.VITE_API_URL + "/disk", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch disks, status: ${res.status}`);
    }

    const data = await res.json();
    const disksPath = DisksSchema.safeParse(data.body);
    if (!disksPath.success) {
      throw new Error("Corrupted data");
    }

    return disksPath.data;
  } catch (error: any) {
    console.error("Error fetching disks:", error);
    return [];
  }
}
