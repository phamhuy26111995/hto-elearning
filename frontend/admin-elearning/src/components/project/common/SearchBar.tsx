"use client";

import { useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Search } from "lucide-react";

export function SearchBar({ onSearch }: { onSearch: (query: string) => void }) {
  const [query, setQuery] = useState("");

  const handleSearch = () => {
    onSearch(query);
  };

  return (
    <div className="flex w-full max-w-md mx-auto">
      <Input
        type="text"
        placeholder="Tìm khóa học..."
        value={query}
        onChange={(e) => setQuery(e.target.value)}
        className="rounded-r-none"
      />
      <Button
        onClick={handleSearch}
        className="rounded-l-none"
        variant="outline"
      >
        <Search className=" h-4 w-4" />
        Tìm kiếm
      </Button>
    </div>
  );
}
