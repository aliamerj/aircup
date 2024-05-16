import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Disks } from "@/schema/disk";
import { useLoaderData } from "react-router-dom";
import { useState } from "react";
import NewDiskInput from "../newDiskInput/NewDiskInput";
export const SelectDisk = () => {
  const [paths, __addNewFolder] = useState(useLoaderData() as Disks);
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

          <NewDiskInput />
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
