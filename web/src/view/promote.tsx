import * as React from "react";
import {ThemedProps} from "./themes";
import {ContainerProps} from "./containers";

export interface CardViewProps extends ThemedProps {
    title: string
    description: string
    onClick?: (e: React.MouseEvent) => void
}

export class CardView extends React.Component<CardViewProps> {
    render() {
        return (
            <div className={"promote-card-view-wrapper"}>
                <div
                    className={`promote-card-view ${this.props.theme}`}
                    onClick={this.props.onClick}>
                    <div className={"title"}>
                        {this.props.title}
                    </div>
                    <div className={"description"}>
                        {this.props.description}
                    </div>
                </div>
            </div>
        )
    }
}

export interface PageTitleProps extends ThemedProps {
    text: string
}

export class PageTitleView extends React.Component<PageTitleProps> {
    render() {
        return (
            <div
                className={`promote-page-title ${this.props.theme}`}>
                {this.props.text}
            </div>
        )
    }
}

export interface PageContainerProps extends ThemedProps {
    title: PageTitleProps
    cards: Array<CardViewProps>
    background: string
}

export class PageContainer extends React.Component<PageContainerProps> {
    render() {
        const cards = this.props.cards.map(props => {
            return <CardView title={props.title} description={props.description} theme={props.theme}/>
        })

        return (
            <div className={`promote-page-container ${this.props.theme} ${this.props.background}`}>
                <PageTitleView text={this.props.title.text} theme={this.props.title.theme}/>
                <div className={"promote-card-container"}>{cards}</div>
            </div>
        )
    }
}

export interface LandingPageInfo {
    title: string
    theme: string
    cards: Array<CardViewProps>
}

export interface LandingProps extends ContainerProps {
    appTitle: string
    appDeveloper: string
    management: LandingPageInfo
    qna: LandingPageInfo
    info: LandingPageInfo
}

export interface LandingState {
    touchStart: number
    direction: number
    touchControl: boolean
    active: number
    allow: boolean
}

export class Landing extends React.Component<LandingProps, LandingState> {
    getInitialState() {
        return {
            touchStart: 0,
            direction: 0,
            touchControl: false,
            active: 0,
            allow: false
        };
    }

    handleMouseDown() {
        this.setState({
            active: this.props.children.length - 1
        });
    }

    moveDown() {
        let active = this.state.active + 1
        if (this.state.active === this.props.children.length - 1) {
            active = this.props.children.length - 1
        }

        this.setState({
            allow: true,
            active: active
        })
    }

    moveUp() {
        this.setState({
            allow: true,
            active: this.state.active === 0 ? 0 : (this.state.active - 1)
        })
    }

    moveTo(n: number) {
        this.setState({
            allow: true,
            active: n
        })
    }

    shouldComponentUpdate(nextProps: LandingProps, nextState: LandingState) {
        return !this.state.allow && nextState.allow
    }

    handleOnWheel(e: React.WheelEvent) {
        if (e.deltaY > 0 && !this.state.allow && this.state.active !== this.props.children.length - 1) {
            this.moveDown()
        }
        if (e.deltaY < 0 && !this.state.allow && this.state.active !== 0) {
            this.moveUp()
        }
    }

    handleTouchStart(e: React.TouchEvent) {
        if (this.state.touchControl) {
            return
        }

        this.setState({
            touchStart: e.touches[0].screenY,
        })
    }

    handleTouchMove(e: React.TouchEvent) {
        if (this.state.touchControl) {
            return
        }
        console.log('touch move')
        let direction = e.touches[0].screenY < this.state.touchStart ? 'down' : 'up'

        if (this.atHead() && direction === 'up') {
            console.log('lala');
            return
        }
        if (this.atEnd() && direction === 'down') {
            console.log('haha');
            return
        }

        if (direction === 'down') {
            this.setState({direction: 1})
        } else {
            this.setState({direction: -1})
        }
    }

    handleTouchEnd(e: React.TouchEvent) {
        if (this.state.touchControl) {
            return
        }

        if (this.state.direction === 1 && !this.atEnd()) {
            this.moveDown()
            this.setState({touchControl: true})
        }
        if (this.state.direction === -1 && !this.atHead()) {
            this.moveUp()
            this.setState({touchControl: true})
        }
    }

    atHead() {
        return this.state.active === 0
    }

    atEnd() {
        return this.state.active === this.props.children.length - 1
    }

    handleLinkClick(n: number) {
        this.moveTo(n)
    }

    handleReset() {
        this.setState({
            allow: false,
            direction: 0,
            touchControl: false
        })
    }

    render() {
        return (
            <div className={"promote-root"}
                 onWheel={ this.handleOnWheel }
                 onTouchStart={ this.handleTouchStart }
                 onTouchMove={ this.handleTouchMove }
                 onTouchEnd={ this.handleTouchEnd }>
                <div className={"promote-app-developer"}>{this.props.appDeveloper}</div>
                <div className={"promote-app-title"}>{this.props.appTitle}</div>
                <PageContainer
                    title={{text: this.props.management.title, theme: this.props.management.theme}}
                    background={"promote-page-task-management"}
                    theme={this.props.management.theme}
                    cards={this.props.management.cards}/>
                <PageContainer
                    title={{text: this.props.qna.title, theme: this.props.qna.theme}}
                    background={"promote-page-qna"}
                    theme={this.props.qna.theme}
                    cards={this.props.qna.cards}/>
                <PageContainer
                    title={{text: this.props.info.title, theme: this.props.info.theme}}
                    background={"promote-page-info"}
                    theme={this.props.info.theme}
                    cards={this.props.info.cards}/>
            </div>
        )
    }
}