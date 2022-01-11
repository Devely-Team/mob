import subprocess

def print_message(message):
    print("")
    print(message)


def run_shell_command(command, silent=True):
    process = subprocess.Popen(command, shell=True, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    while True:
        nextline = process.stdout.readline()
        if nextline == b'' and process.poll() is not None:
            break
        if not silent:
            print(nextline.decode('utf-8').strip())
            
    exitCode = process.returncode

    if (exitCode != 0):
        print("")
        print(process.stderr.read().decode('utf-8'))
    return exitCode

def show_status(status):
    if status == 0:
        print("")
        print("✅ Done!")
    else:
        print("")
        print("❌ Error!")
        print("")
        exit(1)
