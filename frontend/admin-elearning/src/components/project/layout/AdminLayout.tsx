import { BookMarked, UsersRound } from "lucide-react";
import { JSX, useState } from "react";
import { Outlet } from "react-router";
import { Button } from "../../ui/button";
import { Separator } from "../../ui/separator";
import Item from "./components/MenuItem";

export interface MenuItem {
  itemKey: string;
  title: string;
  icon: JSX.Element;
  path: string;
}

const menu: MenuItem[] = [
  {
    itemKey: "students",
    title: "Quản lý học sinh",
    icon: <UsersRound />,
    path: "/students",
  },
  {
    itemKey: "courses",
    title: "Quản lý khóa học",
    icon: <BookMarked />,
    path: "/courses",
  },
];

const AdminLayout = () => {
  const [sidebarOpen, setSidebarOpen] = useState(true);

  return (
    <div className="grid grid-cols-[auto_1fr] h-screen overflow-hidden">
      <aside
        className={`
          ${sidebarOpen ? "w-64" : "w-16"}
          bg-gray-800 text-white flex flex-col
          transition-all duration-300 overflow-hidden
        `}
      >
        <div className="flex items-center justify-between p-4">
          <span className="text-lg font-bold whitespace-nowrap">
            {sidebarOpen ? "Admin" : ""}
          </span>
          <button
            onClick={() => setSidebarOpen(!sidebarOpen)}
            className="p-2 hover:bg-gray-700 rounded"
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-6 w-6"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>
        </div>
        <nav className="flex-1 flex flex-col justify-between">
          <div>
            <ul>
              {menu.map((item) => (
                <Item sidebarOpen={sidebarOpen} item={item} />
              ))}
            </ul>
          </div>
          <div>
            <Separator />
            <div className="p-4">
              <Button className="w-full text-left py-2 hover:bg-gray-700 cursor-pointer">
                Log out
              </Button>
            </div>
          </div>
        </nav>
      </aside>

      <div className="grid grid-rows-[auto_1fr] h-screen">
        <header className="h-16 bg-white shadow flex items-center px-6 justify-between">
          <h1 className="text-2xl font-semibold">Dashboard</h1>
          <button className="px-4 py-2 bg-blue-600 text-white rounded">
            New Item
          </button>
        </header>

        <main className="overflow-auto bg-gray-100 p-6">
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default AdminLayout;
