// TODO : ADD MORE extensions
import {
  FaFile,
  FaFilePdf,
  FaFileWord,
  FaFileExcel,
  FaFilePowerpoint,
  FaFileVideo,
  FaFileAudio,
  FaFileArchive,
  FaFileCode,
  FaFileAlt,
  FaFolder,
  FaFileCsv,
  FaImage,
} from "react-icons/fa";
import { MdFolderOff } from "react-icons/md";
import { AiFillFileMarkdown } from "react-icons/ai";
import { DiJavascript1, DiPython, DiJava, DiRuby, DiPhp } from "react-icons/di";
import {
  BsFiletypeSvg,
  BsFiletypePng,
  BsFiletypeJpg,
  BsFiletypeGif,
  BsFiletypeBmp,
  BsFiletypeMp4,
  BsFiletypeMp3,
  BsFiletypePsd,
} from "react-icons/bs";
import { CiFileOff } from "react-icons/ci";
export const getIconByExtension = (extension: string, ishiding: boolean) => {
  switch (extension.toLowerCase()) {
    case "pdf":
      return <FaFilePdf size={50} className="text-primary" />;
    case "doc":
    case "docx":
      return <FaFileWord size={50} className="text-primary" />;
    case "xls":
    case "xlsx":
      return <FaFileExcel size={50} className="text-primary" />;
    case "ppt":
    case "pptx":
      return <FaFilePowerpoint size={50} className="text-primary" />;
    case "jpg":
      return <BsFiletypeJpg size={50} className="text-primary" />;
    case "jpeg":
      return <FaImage size={50} className="text-primary" />;
    case "png":
      return <BsFiletypePng size={50} className="text-primary" />;
    case "gif":
      return <BsFiletypeGif size={50} className="text-primary" />;
    case "bmp":
      return <BsFiletypeBmp size={50} className="text-primary" />;
    case "svg":
      return <BsFiletypeSvg size={50} className="text-primary" />;
    case "webp":
      return <FaImage size={50} className="text-primary" />;
    case "mp4":
      return <BsFiletypeMp4 size={50} className="text-primary" />;
    case "psd":
      return <BsFiletypePsd size={50} className="text-primary" />;
    case "mkv":
    case "mov":
    case "avi":
    case "flv":
    case "wmv":
      return <FaFileVideo size={50} className="text-primary" />;
    case "mp3":
      return <BsFiletypeMp3 size={50} className="text-primary" />;
    case "wav":
    case "aac":
    case "flac":
    case "ogg":
      return <FaFileAudio size={50} className="text-primary" />;
    case "zip":
    case "rar":
    case "7z":
    case "tar":
    case "gz":
      return <FaFileArchive size={50} className="text-primary" />;
    case "html":
    case "css":
      return <FaFileCode size={50} className="text-primary" />;
    case "js":
      return <DiJavascript1 size={50} className="text-primary" />;
    case "json":
    case "xml":
      return <FaFileCode size={50} className="text-primary" />;
    case "csv":
      return <FaFileCsv size={50} className="text-primary" />;
    case "md":
      return <AiFillFileMarkdown size={50} className="text-primary" />;
    case "py":
      return <DiPython size={50} className="text-primary" />;
    case "java":
      return <DiJava size={50} className="text-primary" />;
    case "rb":
      return <DiRuby size={50} className="text-primary" />;
    case "php":
      return <DiPhp size={50} className="text-primary" />;
    case "txt":
      return <FaFileAlt size={50} className="text-primary" />;
    case "folder":
      return ishiding ? (
        <MdFolderOff size={50} className="text-primary" />
      ) : (
        <FaFolder size={50} className="text-primary" />
      );
    default:
      return ishiding ? (
        <CiFileOff size={50} className="text-primary" />
      ) : (
        <FaFile size={50} className="text-primary" />
      );
  }
};
