from PyInquirer import prompt
from cli.utils import run_shell_command

questions = [{
    "type": "input",
    'qmark': '✨',
    "name": "title",
    "message": "Titulo:"
}, {
    "type": "input",
    'qmark': '✨',
    "name": "description",
    "message": "Descricao:"
}, {
    "type": "list",
    'qmark': '✨',
    "name": "type",
    "message": "tipo:",
    "choices": ["Added", "Changed", "Deprecated", "Removed", "Fixed", "Security"]
}, {
    "type": "input",
    'qmark': '✨',
    "name": "clickup",
    "message": "ClickUp Task ID:"
}, {
    "type": "checkbox",
    'qmark': '✨',
    "name": "category",
    "message": "Categoria:",
    "choices": [
            {'name': '- Bug fix (Correção de um problema sem efeito colateral)'},
            {'name': '- New feature (Adiciona uma nova funcionalidade sem efeito colateral no código)'},
            {'name': '- Breaking change (Correção ou recurso que faria com que alguma funcionalidade existente não funcionasse como esperado)'},
            {'name': '- Code refactor (Melhorias de código)'}
    ]
}, {
    "type": "input",
    'qmark': '✨',
    "name": "module",
    "message": "Modulo:"
}]


def get_template(description, categories):
    template = f'''
# Descrição

{description}

## Qual sua feature flag para essa PR? 

N/A

## Tipo de mudança

'''

    for category in categories:
        template += category

    template += '''

## Como deve ser testado?

N/A

## Prints da tela ou vídeo:

N/A

## Outras observações:

N/A
'''
    return template


def run():
    answers = prompt(questions)
    title = answers.get("title")
    description = answers.get("description")
    type = answers.get("type")
    clickup = answers.get("clickup")
    module = answers.get("module")
    categories = answers.get("category")
    pull_request_title = f"{type}({module}):[{clickup}] {title}"

    template = get_template(description, categories)
    run_shell_command(
        f"gh pr create --title {pull_request_title} --base develop --body {template}", False)
