#include <windows.h>
#include <winnt.h>
#include <winuser.h>

LRESULT CALLBACK WindowProcedure(HWND, UINT, WPARAM, LPARAM);
void AddControls(HWND);

int WINAPI WinMain(HINSTANCE hInst, HINSTANCE hPrevInst, LPSTR args, int ncmdshow) {
    WNDCLASSW wc = {0};

    wc.hbrBackground = (HBRUSH)COLOR_WINDOW;
    wc.hCursor = LoadCursor(NULL, IDC_ARROW);
    wc.hInstance = hInst;
    wc.lpszClassName = L"NoteWindowClass";
    wc.lpfnWndProc = WindowProcedure;

    if (!RegisterClassW(&wc)) {return 1;}

    CreateWindowW(
            L"NoteWindowClass",
            L"Notes",
            WS_OVERLAPPEDWINDOW | WS_VISIBLE,
            100, // x poisition pixels
            100, // y position pixels
            500, // x size pixels
            500, // y size pixels
            NULL, NULL, NULL, NULL
    );

    MSG msg = {0};
    
    while (GetMessage(&msg, NULL, 0, 0)) {
        TranslateMessage(&msg);
        DispatchMessage(&msg);
    }

    return 0;
}

LRESULT CALLBACK WindowProcedure(HWND hWnd, UINT msg, WPARAM wp, LPARAM lp) {
    switch (msg) {
        case WM_CREATE:
            AddControls(hWnd);
            break;
        case WM_DESTROY:
            PostQuitMessage(0);
            return 0;
            break;
        default:
            return DefWindowProcW(hWnd, msg, wp, lp);
    }
}

void AddControls(HWND hWnd) {
    CreateWindowW(
            L"Edit", 
            L"Enter text here.",
            WS_VISIBLE | WS_CHILD | WS_BORDER | ES_MULTILINE,
            0,
            0,
            500,
            500,
            hWnd, NULL, NULL, NULL
        );
}

