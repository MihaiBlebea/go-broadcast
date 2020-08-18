class SocialFooter {
    facebookUrl = null
    linkedinUrl = null
    twitterUrl  = null
    elementId   = null

    constructor(data) {
        if (!data.hasOwnProperty('linkedin')) {
            throw Error('Linkedin share url is missing')
        }

        if (!data.hasOwnProperty('facebook')) {
            throw Error('Facebook share url is missing')
        }

        if (!data.hasOwnProperty('twitter')) {
            throw Error('Twitter share url is missing')
        }

        if (!data.hasOwnProperty('elementId')) {
            throw Error('ElementId is missing')
        }

        this.facebookUrl = data['facebook']
        this.linkedinUrl = data['linkedin']
        this.twitter     = data['twitter']
        this.elementId   = data['elementId']
    }

    scrolling() {
        if (this._scrollPercentage() > 8 && this._scrollPercentage() < 95) {
            if (this._isTriggered() === false) {
                this._attachFooter()
            }
        } else {
            if (this._isTriggered() === true) {
                this._deattachFooter()
            }
        }
    }

    _buildFooter() {
        let wrapper = window.document.createElement('DIV')
        wrapper.classList.add('footer-social')
        wrapper.id = this.elementId

        // Label
        let label = window.document.createElement('SPAN')
        label.classList.add('mr-2')
        label.innerHTML = 'Share on'

        // Linkedin
        let linkedinAnchor = window.document.createElement('A')
        linkedinAnchor.classList.add('no-underline')
        linkedinAnchor.href = linkedinUrl
        linkedinAnchor.target = '_blank'

        let linkedinIcon = window.document.createElement('I')
        linkedinIcon.className = 'fa fa-linkedin-square mr-2'
        linkedinAnchor.appendChild(linkedinIcon)

        // Facebook
        let facebookAnchor = window.document.createElement('A')
        facebookAnchor.classList.add('no-underline')
        facebookAnchor.href = facebookUrl
        facebookAnchor.target = '_blank'

        let facebookIcon = window.document.createElement('I')
        facebookIcon.className = 'fa fa-facebook-square mr-2'
        facebookAnchor.appendChild(facebookIcon)

        // Twitter
        let twitterAnchor = window.document.createElement('A')
        twitterAnchor.classList.add('no-underline')
        twitterAnchor.href = twitterUrl
        twitterAnchor.target = '_blank'

        let twitterIcon = window.document.createElement('I')
        twitterIcon.className = 'fa fa-twitter-square mr-2'
        twitterAnchor.appendChild(twitterIcon)

        wrapper.appendChild(label)
        wrapper.appendChild(linkedinAnchor)
        wrapper.appendChild(facebookAnchor)
        wrapper.appendChild(twitterAnchor)

        return wrapper
    }

    _attachFooter() {
        document.body.appendChild(
            this._buildFooter()
        )
    }

    _deattachFooter() {
        document.body.removeChild(
            document.getElementById(this.elementId)
        )
    }

    _scrollPercentage() {
        let current = document.body.scrollTop || document.documentElement.scrollTop
        let total = document.body.scrollHeight
        let screen = document.body.clientHeight

        return ((current + screen) / total) * 100
    }

    _isTriggered() {
        return document.getElementById(this.elementId) !== null
    }
}