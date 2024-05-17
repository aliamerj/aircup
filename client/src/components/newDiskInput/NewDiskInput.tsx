import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { CircleCheckBig, CircleHelp, CircleX, SaveIcon } from "lucide-react";
import { useState } from "react";
import { toast } from "../ui/use-toast";

function NewDiskInput({ addNewDir }: { addNewDir: (newPath: string) => void }) {
  const [isAvaliable, setIsAvalible] = useState<null | string>(null);
  const [sendRequest, setSendRequest] = useState<
    "NOT" | "START" | "FAILED" | "SUCCESS"
  >("NOT");
  const checkDirPath = async (event: React.FocusEvent<HTMLInputElement>) => {
    event.preventDefault();
    if (event.target.value.length === 0) return;
    try {
      const res = await fetch(
        import.meta.env.VITE_API_URL + "/disk/check?path=" + event.target.value,
        {
          method: "GET",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      if (!res.ok) {
        return setIsAvalible("");
      }
      return setIsAvalible(event.target.value);
    } catch (error) {
      setIsAvalible("");
    }
  };
  const savePath = async (event: React.MouseEvent<HTMLButtonElement>) => {
    event.preventDefault();
    setSendRequest("START");
    if (!isAvaliable) return;
    try {
      const response = await fetch(
        import.meta.env.VITE_API_URL + "/disk/save?path=" + isAvaliable,
        {
          method: "GET",
          credentials: "include",
          headers: {
            "Content-Type": "application/json",
          },
        },
      );
      const result = await response.json();
      if (!response.ok) {
        throw new Error(result.message);
      }

      addNewDir(isAvaliable);
      setSendRequest("SUCCESS");
    } catch (error: any) {
      console.log({ error });
      setSendRequest("FAILED");
      return toast({
        variant: "destructive",
        title: "Uh oh! Something went wrong.",
        description: error.message,
      });
    }
  };
  return (
    <Dialog open={sendRequest === "SUCCESS" ? false : undefined}>
      <DialogTrigger asChild>
        <Button size="lg" className="w-full rounded-md">
          Add new Folder/Disk
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-md">
        <DialogHeader>
          <DialogTitle>Add New Folder</DialogTitle>
          <DialogDescription>
            Please enter the full path of the folder you want to share. Make
            sure the path is correct and accessible.{" "}
          </DialogDescription>
        </DialogHeader>
        <div className="flex items-center space-x-2">
          <div className="grid flex-1 gap-2">
            <Label htmlFor="link" className="sr-only">
              Link
            </Label>
            <Input
              id="link"
              placeholder="Enter the full path of the folder"
              onBlur={checkDirPath}
              onChange={() => setIsAvalible(null)}
            />
          </div>
          {isAvaliable === null ? (
            <CircleHelp className="h-7 w-7" />
          ) : isAvaliable === "" ? (
            <CircleX className="h-7 w-7" />
          ) : (
            <CircleCheckBig className="h-7 w-7" />
          )}
        </div>
        <DialogFooter className="w-full items-center gap-2 sm:justify-start">
          <DialogClose asChild>
            <Button
              type="button"
              variant="secondary"
              className="w-full sm:w-1/2"
            >
              Close
            </Button>
          </DialogClose>
          {isAvaliable && (
            <Button
              type="submit"
              size="sm"
              className="w-full px-3 sm:w-1/2"
              onClick={savePath}
            >
              <SaveIcon className="h-8 w-7" />
            </Button>
          )}
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}

export default NewDiskInput;
