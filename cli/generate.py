from cli.utils import run_shell_command, print_message, show_status

def run():
    print_message("📦 Generating...")

    print_message("🚀 Running flutter pub get...")
    status = run_shell_command("flutter pub get")
    show_status(status)

    print_message("🚀 Generating assets...")
    status = run_shell_command("fluttergen")
    show_status(status)

    print_message("🚀 Generating i18n...")
    status = run_shell_command("flutter pub run fast_i18n")
    show_status(status)

    print_message("🚀 Generating xcode project...")
    status = run_shell_command("cd ios && xcodegen generate && pod install")
    show_status(status)