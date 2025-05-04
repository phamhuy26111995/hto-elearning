import { Switch } from "@/components/ui/switch";
import { useDarkMode } from "@/hooks/useDarkMode";
import { Label } from "@radix-ui/react-dropdown-menu";
import { Moon, Sun } from "lucide-react";

export default function SwitchDarkMode() {
  const [isDark, setIsDark] = useDarkMode();

  function toggleDarkMode() {
    setIsDark(!isDark);
  }

  return (
    <div className="flex items-center space-x-2 p-5">
      <Switch id="airplane-mode" onCheckedChange={toggleDarkMode} />
      <Label>{isDark ? <Moon /> : <Sun />}</Label>
    </div>
  );
}
