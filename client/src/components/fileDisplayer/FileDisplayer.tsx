import { format } from "date-fns";
import { getIconByExtension } from "@/lib/getIconByExtension";

export const FileDisplayer = ({
  fileName,
  extension,
  size,
  modifiedTime,
}: {
  fileName: string;
  extension: string | null;
  size: string;
  modifiedTime: string;
}) => {
  return (
    <div className="flex h-48 w-56 transform flex-col items-center justify-center bg-secondary p-4 transition-transform hover:scale-105 hover:shadow-xl">
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
