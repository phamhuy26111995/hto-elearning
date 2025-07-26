import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from '@/components/ui/tooltip';
import { cn } from '@/lib/utils';
import React from 'react';

interface ButtonTooltipProps {
  icon: React.ReactNode;
  tooltip: string;
  onClick?: () => void;
  className?: string;
}

export default function ButtonTooltip({
  icon,
  tooltip,
  onClick,
  className,
}: ButtonTooltipProps) {
  return (
    <TooltipProvider>
      <Tooltip>
        <TooltipTrigger asChild>
          <button
            className={cn('hover:cursor-pointer', className)}
            onClick={onClick}
            type="button"
          >
            {icon}
          </button>
        </TooltipTrigger>
        <TooltipContent>
          <span>{tooltip}</span>
        </TooltipContent>
      </Tooltip>
    </TooltipProvider>
  );
}
