import { Button } from "@/components/ui/button"
import { Plus } from "lucide-react"

type FloatingButtonProps = React.ComponentProps<typeof Button> & {
  content?: React.ReactNode
}

export default function FloatingButton({ content, ...props} : FloatingButtonProps) {
  return (
    <Button
     {...props}
      className="fixed bottom-6 right-8 rounded-full shadow-lg w-14 h-14 p-0 text-white bg-blue-600 hover:bg-blue-700"
      size="icon"
    >
      {content || <Plus className="size-6" />}
    </Button>
  )
}