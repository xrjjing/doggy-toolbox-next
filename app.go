package main

import (
    "context"
    "fmt"
    "strings"
)

type App struct {
    ctx context.Context
}

type AppInfo struct {
    Name   string `json:"name"`
    Stack  string `json:"stack"`
    Status string `json:"status"`
}

func NewApp() *App {
    return &App{}
}

func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
}

func (a *App) Greet(name string) string {
    finalName := strings.TrimSpace(name)
    if finalName == "" {
        finalName = "Doggy"
    }
    return fmt.Sprintf("Hello %s, doggy-toolbox-next is ready.", finalName)
}

func (a *App) GetAppInfo() AppInfo {
    return AppInfo{
        Name:   "doggy-toolbox-next",
        Stack:  "Go + Wails + Vue 3 + TypeScript",
        Status: "可运行白板基线已就绪",
    }
}
