from cli.utils import run_shell_command, print_message, show_status


def run(dev=False):
    print_message("📦 Installing dependencies...")

    print_message("🔧 Installing flutter_gen dependencies...")
    status = run_shell_command("dart pub global activate flutter_gen")
    show_status(status)

    print_message("🔧 Installing xcodegen...")
    status = run_shell_command("brew install xcodegen")
    show_status(status)

    print_message("🚀 Running flutter pub get...")
    status = run_shell_command("flutter pub get")
    show_status(status)

    if dev:
        print_message("🔧 Installing mason...")
        status = run_shell_command(
            "dart pub global activate mason_cli && mason get")
        show_status(status)

        print_message("🔧 Installing dev dependencies (Github CLI)...")
        status = run_shell_command("brew install gh")
        show_status(status)
