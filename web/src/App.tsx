import React from 'react';
import './css/fonts.css';
import './css/app.css';
import './css/dark.css';
import './css/promote.css';
import * as Promote from "./view/promote";

function App() {
    return (
        <Promote.Landing
            appDeveloper={"Projector Solutions"}
            appTitle={"Pharaon"}
            management={{
                title: "Комплексное решение для управления задачами",
                theme: "light",
                cards: [{
                    title: "Разработка",
                    description: "Организуйте работу команды разработчиков по методологии AGILE: планируйте спринты, контролируйте ход работы над задачами на доске и оценивайте нагрузку в очередях.",
                    theme: "dark"
                }, {
                    title: "Маркетинг",
                    description: "Соберите дашборд подготовки рекламной кампании, запустите автоматический анализ для формирования графики и статистики. Учет времени и комментарии сделают работу  более прозрачной.",
                    theme: "dark"
                }, {
                    title: "Workflow",
                    description: "Настраивайте статусы задач и переходы по ним, чтобы всей команде было понятно, на какой стадии находится проект. Оптимизируйте процессы, получая информацию обо всех задачах на дашбордах",
                    theme: "dark"
                }, {
                    title: "Инфраструктура",
                    description: "Решайте задачи инвентаризации, закупки или обслуживания оборудования. Ведите учет расходов и контроль SLA проектов.",
                    theme: "dark"
                }, {
                    title: "BOX edition",
                    description: "Обеспечьте своим проектам макисмальный уровень защиты и надежности - разверните локальную версию Pharaon на своем сервере. Доступ ко всем данным будет иметь только команда.",
                    theme: "dark"
                }, {
                    title: "API",
                    description: "Используйте API сервиса для интеграции ботов, мессенджеров и других систем.\nЗаходите в маркетплейс и выбирайте нужные дополнения от других разработчиков.",
                    theme: "dark"
                }]
            }}
            qna={{
                title: "Вопросы и ответы",
                theme: "light",
                cards: []
            }}
            info={{
                title: "Информация",
                theme: "dark",
                cards: []
            }}
            children={[]}/>
    );
}

export default App;

