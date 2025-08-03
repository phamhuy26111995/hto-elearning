import React, { useMemo } from "react";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationPrevious,
  PaginationNext,
  PaginationEllipsis,
} from "@/components/ui/pagination";

interface PagingProps {
  totalRows: number;
  rowsPerPage: number;
  currentPage: number;
  onChange: (pageNumber: number, rowsPerPage: number) => void;
}

const Paging: React.FC<PagingProps> = ({
  totalRows,
  rowsPerPage,
  currentPage,
  onChange,
}) => {
  const totalPages = useMemo(() => Math.ceil(totalRows / rowsPerPage), [totalRows, rowsPerPage]);

  const handlePageClick = (page: number) => {
    if (page >= 1 && page <= totalPages && page !== currentPage) {
      onChange(page, rowsPerPage);
    }
  };

  const renderPageNumbers = () => {
    const pages = [];

    if (totalPages <= 5) {
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      // Hiển thị tối đa 5 trang (ví dụ: 1 ... 4 5 6 ... 10)
      if (currentPage <= 3) {
        pages.push(1, 2, 3, 4, "ellipsis", totalPages);
      } else if (currentPage >= totalPages - 2) {
        pages.push(1, "ellipsis", totalPages - 3, totalPages - 2, totalPages - 1, totalPages);
      } else {
        pages.push(1, "ellipsis", currentPage - 1, currentPage, currentPage + 1, "ellipsis", totalPages);
      }
    }

    return pages.map((page, idx) => {
      if (page === "ellipsis") {
        return (
          <PaginationItem key={`ellipsis-${idx}`}>
            <PaginationEllipsis />
          </PaginationItem>
        );
      }

      return (
        <PaginationItem key={page}>
          <PaginationLink
            href="#"
            isActive={page === currentPage}
            onClick={(e) => {
              e.preventDefault();
              handlePageClick(page as number);
            }}
          >
            {page}
          </PaginationLink>
        </PaginationItem>
      );
    });
  };

  return (
    <Pagination>
      <PaginationContent>
        {currentPage > 1 && (
          <PaginationItem>
            <PaginationPrevious
              href="#"
              onClick={(e) => {
                e.preventDefault();
                handlePageClick(currentPage - 1);
              }}
            />
          </PaginationItem>
        )}

        {renderPageNumbers()}

       {currentPage < totalPages && (
          <PaginationItem>
            <PaginationNext
              href="#"
              onClick={(e) => {
                e.preventDefault();
                handlePageClick(currentPage + 1);
              }}
            />
          </PaginationItem>
        )}  
      </PaginationContent>
    </Pagination>
  );
};

export default Paging;
