import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useLoaderData } from "react-router-dom";
import { useState } from "react";
import NewDiskInput from "../newDiskInput/NewDiskInput";
import { SavedDiskPaths } from "@/schema/disk_path";
export const SelectDisk = () => {
  const [paths, addNewFolder] = useState<SavedDiskPaths>(
    useLoaderData() as SavedDiskPaths,
  );
  const addNewDir = (newPath: string) => {
    addNewFolder((current) => [...current, { disk_path: newPath }]);
  };
  return (
    <Select>
      <SelectTrigger className="w-1/4">
        <SelectValue placeholder="Select Folder or disk" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Folders/Disks</SelectLabel>
          {paths.map((path) => (
            <SelectItem key={path.disk_path} value={path.disk_path}>
              {path.disk_path}
            </SelectItem>
          ))}

          <NewDiskInput addNewDir={addNewDir} />
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
