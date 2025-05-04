"use client"

import {
    AudioWaveform,
    Bot,
    Command,
    GalleryVerticalEnd,
    Settings2,
    SquareTerminal
} from "lucide-react"
import * as React from "react"

import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarHeader,
    SidebarRail,
} from "@/components/ui/sidebar"
import { NavMain } from "./NavMain"
import { NavUser } from "./NavUser"
import { TeamSwitcher } from "./TeamSwitcher"


const data = {
  user: {
    name: "shadcn",
    email: "m@example.com",
    avatar: "/avatars/shadcn.jpg",
  },
  teams: [
    {
      name: "Acme Inc",
      logo: GalleryVerticalEnd,
      plan: "Enterprise",
    },
    {
      name: "Acme Corp.",
      logo: AudioWaveform,
      plan: "Startup",
    },
    {
      name: "Evil Corp.",
      logo: Command,
      plan: "Free",
    },
  ],
  navMain: [
    {
      title: "Quản lý khóa học",
      url: "#",
      icon: SquareTerminal,
      isActive: true,
      items: [
        {
          title: "Xem danh sách khóa học",
          url: "/courses",
        },
        {
          title: "Tạo mới khóa học",
          url: "/course/new",
        },
      ],
    },
    {
      title: "Quản lý học sinh",
      url: "#",
      icon: Bot,
      items: [
        {
          title: "Xem danh sách học sinh",
          url: "/students",
        },
        {
          title: "Tạo mới học sinh",
          url: "/student/new",
        },
      ],
    },

    {
      title: "Cấu hình",
      url: "#",
      icon: Settings2,
      items: [
        {
          title: "General",
          url: "#",
        },
        {
          title: "Team",
          url: "#",
        },
        {
          title: "Billing",
          url: "#",
        },
        {
          title: "Limits",
          url: "#",
        },
      ],
    },
  ]
}

export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar collapsible="icon" {...props}>
      <SidebarHeader>
        <TeamSwitcher teams={data.teams} />
      </SidebarHeader>
      <SidebarContent>
        <NavMain items={data.navMain} />
      </SidebarContent>
      <SidebarFooter>
        <NavUser user={data.user} />
      </SidebarFooter>
      <SidebarRail />
    </Sidebar>
  )
}
