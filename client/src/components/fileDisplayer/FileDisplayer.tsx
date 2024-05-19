import { format } from "date-fns";
import { getIconByExtension } from "@/lib/getIconByExtension";


export const FileDisplayer = ({
  path,
  fileName,
  extension,
  size,
  modifiedTime,
  setDisk,
}: {
  path: string;
  fileName: string;
  extension: string | null;
  size: string;
  modifiedTime: string;
  setDisk: (path: string) => void;
}) => {
  const openFolder = () => {
    if (extension === "folder") {
      setDisk(`${path}/${fileName}`);
    }
  };

  const isFolder = extension === "folder";
  return (
    <div
      onClick={openFolder}
      className={`flex h-48 w-56 transform flex-col items-center justify-center bg-secondary p-4 transition-transform hover:scale-105 hover:shadow-xl ${
        isFolder ? "cursor-pointer" : ""
      }`}
    >
      {getIconByExtension(extension ?? "folder", fileName.startsWith("."))}
      <span
        className="mt-4 w-full truncate text-center text-lg font-medium text-primary"
        title={fileName}
      >
        {fileName.length > 20 ? fileName.slice(0, 20) + "..." : fileName}
      </span>
      <span className="text-sm text-gray-500">{size}</span>
      <span className="text-sm text-gray-500">
        {format(new Date(modifiedTime), "dd MMM yyyy, hh:mm a")}
      </span>
    </div>
  );
};

