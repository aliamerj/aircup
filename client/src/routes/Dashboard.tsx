import { Button } from "@/components/ui/button";

import { Sidbar } from "@/components/sidebar/Sidebar";
import { Navbar } from "@/components/Navbar/Header";
export function Dashboard() {
  return (
    <div className="grid min-h-screen w-full md:grid-cols-[220px_1fr] lg:grid-cols-[280px_1fr]">
      <Sidbar />
      <div className="flex flex-col">
        <Navbar />
        <main className="flex flex-1 flex-col gap-4 p-4 lg:gap-6 lg:p-6">
          <div className="flex items-center">
            <h1 className="text-lg font-semibold md:text-2xl">My Files</h1>
          </div>
          <div
            className="flex flex-1 items-center justify-center rounded-lg border border-dashed shadow-sm"
            x-chunk="dashboard-02-chunk-1"
          >
            <div className="flex flex-col items-center gap-1 text-center">
              <h3 className="text-2xl font-bold tracking-tight">
                You have no File
              </h3>
              <p className="text-sm text-muted-foreground">
                You can start sharing your files as soon as you add them.
              </p>
              <Button className="mt-4">Add File</Button>
            </div>
          </div>
        </main>
      </div>
    </div>
  );
}
