import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from '@/components/ui/alert-dialog';
import { Info } from 'lucide-react';

type DialogRemoveProps = {
  open: boolean;
  onOpenChange: (open: boolean) => void;
  username?: string;
  onConfirm: () => void;
  title?: string;
  description?: string;
  confirmText?: string;
  cancelText?: string;
};

export default function DialogRemove({
  open,
  onOpenChange,
  username,
  onConfirm,
  title = 'Bạn có chắc muốn xóa mục này?',
  description = 'Việc xóa sẽ không thể hoàn tác. Bạn có chắc chắn muốn tiếp tục?',
  confirmText = 'Có',
  cancelText = 'Không',
}: DialogRemoveProps) {
  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle className="text-center">
            <div className="flex justify-center mb-5">
              <Info size={40} className="text-red-500" />
            </div>
            <p>{title}</p>
            {username && (
              <p className="font-bold text-xl text-red-500">{username}</p>
            )}
          </AlertDialogTitle>
          <AlertDialogDescription className="text-center">
            {description}
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel>{cancelText}</AlertDialogCancel>
          <AlertDialogAction
            onClick={() => {
              onConfirm();
              onOpenChange(false);
            }}
          >
            {confirmText}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  );
}
