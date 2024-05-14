import { Moon, Sun } from "lucide-react";

import { Button } from "@/components/ui/button";
import { useTheme } from "../theme/theme-hook";

export function ModeToggle() {
  const { setTheme, theme } = useTheme();
  const changeTheme = () => {
    if (theme !== "light") return setTheme("light");
    return setTheme("dark");
  };

  return (
    <Button variant="outline" size="icon" onClick={() => changeTheme()}>
      <Sun className="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0" />
      <Moon className="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100" />
      <span className="sr-only">Toggle theme</span>
    </Button>
  );
}
