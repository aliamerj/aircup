import { AlertTriangle } from "lucide-react";

const ErrorDisplay = ({ message }: { message: string }) => {
  return (
    <div className="flex min-h-screen items-center justify-center bg-gradient-to-r from-background to-transparent">
      <div className="transform rounded-lg bg-muted p-10 text-center shadow-xl transition duration-300 hover:scale-105">
        <div className="animate-spin-slow mb-4">
          <AlertTriangle className="mx-auto h-16 w-16 text-red-500" />
        </div>
        <h1 className="mb-2 text-2xl font-semibold text-primary">
          Oops! Something went wrong...
        </h1>
        <p className="text-gray-600">{message}</p>
      </div>
    </div>
  );
};

export default ErrorDisplay;
