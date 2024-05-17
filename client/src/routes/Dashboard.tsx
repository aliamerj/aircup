import { Sidbar } from "@/components/sidebar/Sidebar";
import { SelectDisk } from "@/components/selectDisk/SelectDisk";
import { Navbar } from "@/components/Navbar/Navbar";
import { useState, useTransition } from "react";
import { Disk } from "@/schema/disk";
import { loadDisk } from "@/api-handler/disk-actions";
import { Loader } from "@/components/loader/Loader";

import { FileDisplayer } from "@/components/fileDisplayer/FileDisplayer";

export function Dashboard() {
  const [isPending, startTransition] = useTransition();
  const [disk, setDisk] = useState<Disk | null>(null);
  const selectedDisk = (path: string) => {
    startTransition(() => {
      const load = async () => {
        const disk = await loadDisk(path);
        setDisk(disk);
      };
      load();
    });
  };
  return (
    <div className="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
      <Sidbar />
      <div className="flex flex-col">
        <Navbar />
        <main className="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-lg font-semibold md:text-2xl">My Files</h1>
              {disk && (
                <div>
                  <p>Total Size: {disk.totalSize}</p>
                  <p>Available Size: {disk.availableSize}</p>
                </div>
              )}
            </div>
            <SelectDisk selectDiskPath={selectedDisk} />
          </div>
          <div className="h-full border border-dashed">
            {isPending ? (
              <Loader />
            ) : disk && disk.contents.length > 0 ? (
              <div className="flex flex-wrap gap-4 p-4">
                {disk.contents.map((opt, i) => (
                  <FileDisplayer
                    key={i}
                    fileName={opt.name}
                    extension={opt.extension}
                    size={opt.size}
                    modifiedTime={opt.modifiedTime}
                  />
                ))}
              </div>
            ) : (
              <div className="flex h-full flex-col items-center justify-center text-center">
                <h3 className="text-2xl font-bold tracking-tight text-gray-800">
                  You have no files in the selected folder
                </h3>
              </div>
            )}
          </div>
        </main>
      </div>
    </div>
  );
}
