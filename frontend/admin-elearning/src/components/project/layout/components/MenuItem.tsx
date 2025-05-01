import React from "react";
import { MenuItem } from "../AdminLayout";

type Props = {
  item: MenuItem;
  sidebarOpen: boolean;
};

export default function Item({ item, sidebarOpen }: Props) {
  return (
    <li
      key={item.itemKey}
      className="px-4 py-2 hover:bg-gray-700 cursor-pointer"
    >
      {sidebarOpen ? (
        <div className="flex items-center gap-2">
          <span>{item.icon}</span>
          <span>{item.title}</span>
        </div>
      ) : (
        <div className="flex justify-center items-center">
          <span>{item.icon}</span>
        </div>
      )}
    </li>
  );
}
