import { Label } from '@/components/ui/label';

export function RequiredLabel({
  htmlFor,
  children,
}: {
  htmlFor: string;
  children: React.ReactNode;
}) {
  return (
    <Label htmlFor={htmlFor}>
      <span className='text-gray-700 font-medium mb-3'>
        {children} (<span className="text-red-500">*</span>)
      </span>
    </Label>
  );
}
