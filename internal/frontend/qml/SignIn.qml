// Copyright (c) 2021 Proton Technologies AG
//
// This file is part of ProtonMail Bridge.
//
// ProtonMail Bridge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ProtonMail Bridge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ProtonMail Bridge.  If not, see <https://www.gnu.org/licenses/>.

import QtQml 2.12
import QtQuick 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.impl 2.12

import Proton 4.0

Item {
    id: root
    property var colorScheme: parent.colorScheme

    function abort() {
        root.loginAbort(usernameTextField.text)
    }

    signal login(string username, string password)
    signal login2FA(string username, string code)
    signal login2Password(string username, string password)
    signal loginAbort(string username)

    implicitHeight: children[0].implicitHeight
    implicitWidth: children[0].implicitWidth

    property var backend
    property var window

    // in case of adding new account this property should be undefined
    property var user
    state: "Page 1"

    onUserChanged: {
        stackLayout.currentIndex = 0
        loginNormalLayout.reset()
        passwordTextField.text = ""
        login2FALayout.reset()
        login2PasswordLayout.reset()
    }

    onLoginAbort: {
        stackLayout.currentIndex = 0
        loginNormalLayout.reset()
        login2FALayout.reset()
        login2PasswordLayout.reset()
    }

    property alias currentIndex: stackLayout.currentIndex

    StackLayout {
        id: stackLayout
        anchors.fill: parent

        function loginFailed() {
            signInButton.loading = false

            usernameTextField.enabled = true
            usernameTextField.error = true

            passwordTextField.enabled = true
            passwordTextField.error = true
        }

        Connections {
            target: user !== undefined ? user : root.backend

            onLoginUsernamePasswordError: {
                console.assert(stackLayout.currentIndex == 0, "Unexpected loginUsernamePasswordError")
                console.assert(signInButton.loading == true, "Unexpected loginUsernamePasswordError")

                stackLayout.loginFailed()
                errorLabel.text = qsTr("Your email and/or password are incorrect")

            }

            onLoginFreeUserError: {
                console.assert(stackLayout.currentIndex == 0, "Unexpected loginFreeUserError")
                stackLayout.loginFailed()
                window.notifyOnlyPaidUsers()
            }
            onLoginConnectionError: {
                if (stackLayout.currentIndex == 0 ) {
                    stackLayout.loginFailed()
                }
                window.notifyConnectionLostWhileLogin()
            }

            onLogin2FARequested: {
                console.assert(stackLayout.currentIndex == 0, "Unexpected login2FARequested")

                stackLayout.currentIndex = 1
            }
            onLogin2FAError: {
                console.assert(stackLayout.currentIndex == 1, "Unexpected login2FAError")

                twoFAButton.loading = false

                twoFactorPasswordTextField.enabled = true
                twoFactorPasswordTextField.error = true
                twoFactorPasswordTextField.assistiveText = qsTr("Your code is incorrect")
            }
            onLogin2FAErrorAbort: {
                console.assert(stackLayout.currentIndex == 1, "Unexpected login2FAErrorAbort")

                stackLayout.currentIndex = 0
                loginNormalLayout.reset()
                login2FALayout.reset()
                login2PasswordLayout.reset()

                errorLabel.text = qsTr("Incorrect login credentials. Please try again.")
                passwordTextField.text = ""
            }

            onLogin2PasswordRequested: {
                console.assert(stackLayout.currentIndex == 0 || stackLayout.currentIndex == 1, "Unexpected login2PasswordRequested")

                stackLayout.currentIndex = 2
            }
            onLogin2PasswordError: {
                console.assert(stackLayout.currentIndex == 2, "Unexpected login2PasswordError")

                secondPasswordButton.loading = false

                secondPasswordTextField.enabled = true
                secondPasswordTextField.error = true
                secondPasswordTextField.assistiveText = qsTr("Your mailbox password is incorrect")
            }
            onLogin2PasswordErrorAbort: {
                console.assert(stackLayout.currentIndex == 2, "Unexpected login2PasswordErrorAbort")

                stackLayout.currentIndex = 0
                loginNormalLayout.reset()
                login2FALayout.reset()
                login2PasswordLayout.reset()

                errorLabel.text = qsTr("Incorrect login credentials. Please try again.")
                passwordTextField.text = ""
            }
        }

        ColumnLayout {
            id: loginNormalLayout

            function reset() {
                signInButton.loading = false

                errorLabel.text = ""

                usernameTextField.enabled = true
                usernameTextField.error = false
                usernameTextField.assistiveText = ""

                passwordTextField.enabled = true
                passwordTextField.error = false
                passwordTextField.assistiveText = ""
            }

            spacing: 0

            ProtonLabel {
                text: qsTr("Sign in")
                Layout.alignment: Qt.AlignHCenter
                Layout.topMargin: 16
                font.weight: ProtonStyle.fontWidth_700
            }

            ProtonLabel {
                id: subTitle
                text: qsTr("Enter your Proton Account details.")
                Layout.alignment: Qt.AlignHCenter
                Layout.topMargin: 8
                color: root.colorScheme.text_weak
                state: "body"
            }

            RowLayout {
                Layout.fillWidth: true
                Layout.topMargin: 36

                spacing: 0
                visible: errorLabel.text.length > 0

                ColorImage {
                    color: root.colorScheme.signal_danger
                    source: "./icons/ic-exclamation-circle-filled.svg"
                }

                ProtonLabel {
                    id: errorLabel
                    Layout.leftMargin: 4
                    color: root.colorScheme.signal_danger

                    font.weight: root.error ? ProtonStyle.fontWidth_600 : ProtonStyle.fontWidth_400
                    state: "caption"
                }
            }

            TextField {
                id: usernameTextField
                label: qsTr("Username or email")

                text: user !== undefined ? user.username : ""

                Layout.fillWidth: true
                Layout.topMargin: 24

                onTextEdited: { // TODO: repeating?
                    if (error || errorLabel.text.length > 0) {
                        errorLabel.text = ""

                        usernameTextField.error = false
                        usernameTextField.assistiveText = ""

                        passwordTextField.error = false
                        passwordTextField.assistiveText = ""
                    }
                }

                onAccepted: passwordTextField.forceActiveFocus()
            }

            TextField {
                id: passwordTextField
                label: qsTr("Password")

                Layout.fillWidth: true
                Layout.topMargin: 8
                echoMode: TextInput.Password

                onTextEdited: {
                    if (error || errorLabel.text.length > 0) {
                        errorLabel.text = ""

                        usernameTextField.error = false
                        usernameTextField.assistiveText = ""

                        passwordTextField.error = false
                        passwordTextField.assistiveText = ""
                    }
                }

                onAccepted: signInButton.checkAndSignIn()
            }

            Button {
                id: signInButton
                text: qsTr("Sign in")

                Layout.fillWidth: true
                Layout.topMargin: 24


                onClicked: checkAndSignIn()

                function checkAndSignIn() {
                    var err = false

                    if (usernameTextField.text.length == 0) {
                        usernameTextField.error = true
                        usernameTextField.assistiveText = qsTr("Enter username or email")
                        err = true
                    } else {
                        usernameTextField.error = false
                        usernameTextField.assistiveText = qsTr("")
                    }

                    if (passwordTextField.text.length == 0) {
                        passwordTextField.error = true
                        passwordTextField.assistiveText = qsTr("Enter password")
                        err = true
                    } else {
                        passwordTextField.error = false
                        passwordTextField.assistiveText = qsTr("")
                    }

                    if (err) {
                        return
                    }

                    usernameTextField.enabled = false
                    passwordTextField.enabled = false

                    enabled = false
                    loading = true

                    if (root.user !== undefined) {
                        root.user.login(usernameTextField.text, passwordTextField.text)
                        return
                    }

                    root.login(usernameTextField.text, passwordTextField.text)
                }
            }

            ProtonLabel {
                textFormat: Text.StyledText
                text: putLink("https://protonmail.com/upgrade", qsTr("Create or upgrade your account"))
                Layout.alignment: Qt.AlignHCenter
                Layout.topMargin: 24
                state: "body"

                onLinkActivated: {
                    Qt.openUrlExternally(link)
                }

            }
        }

        ColumnLayout {
            id: login2FALayout

            function reset() {
                twoFAButton.loading = false

                twoFactorPasswordTextField.enabled = true
                twoFactorPasswordTextField.error = false
                twoFactorPasswordTextField.assistiveText = ""
            }

            spacing: 0

            ProtonLabel {
                text: qsTr("Two-factor authentication")
                Layout.topMargin: 16
                Layout.alignment: Qt.AlignCenter
                font.weight: ProtonStyle.fontWidth_700


            }

            TextField {
                id: twoFactorPasswordTextField
                label: qsTr("Two-factor authentication code")

                Layout.fillWidth: true
                Layout.topMargin: 8 + implicitHeight + 24 + subTitle.implicitHeight

                onTextEdited: {
                    if (error) {
                        twoFactorPasswordTextField.error = false
                        twoFactorPasswordTextField.assistiveText = ""
                    }
                }
            }

            Button {
                id: twoFAButton
                text: loading ? qsTr("Authenticating") : qsTr("Authenticate")

                Layout.fillWidth: true
                Layout.topMargin: 24

                onClicked: {
                    var err = false

                    if (twoFactorPasswordTextField.text.length == 0) {
                        twoFactorPasswordTextField.error = true
                        twoFactorPasswordTextField.assistiveText = qsTr("Enter username or email")
                        err = true
                    } else {
                        twoFactorPasswordTextField.error = false
                        twoFactorPasswordTextField.assistiveText = qsTr("")
                    }

                    if (err) {
                        return
                    }

                    twoFactorPasswordTextField.enabled = false

                    enabled = false
                    loading = true

                    if (root.user !== undefined) {
                        root.user.login2FA(usernameTextField.text, twoFactorPasswordTextField.text)
                        return
                    }

                    root.login2FA(usernameTextField.text, twoFactorPasswordTextField.text)
                }
            }
        }

        ColumnLayout {
            id: login2PasswordLayout

            function reset() {
                secondPasswordButton.loading = false

                secondPasswordTextField.enabled = true
                secondPasswordTextField.error = false
                secondPasswordTextField.assistiveText = ""
            }

            spacing: 0

            ProtonLabel {
                text: qsTr("Unlock your mailbox")
                Layout.topMargin: 16
                Layout.alignment: Qt.AlignCenter
                font.weight: ProtonStyle.fontWidth_700
            }






            TextField {
                id: secondPasswordTextField
                label: qsTr("Mailbox password")

                Layout.fillWidth: true
                Layout.topMargin: 8 + implicitHeight + 24 + subTitle.implicitHeight
                echoMode: TextInput.Password

                onTextEdited: {
                    if (error) {
                        secondPasswordTextField.error = false
                        secondPasswordTextField.assistiveText = ""
                    }
                }
            }

            Button {
                id: secondPasswordButton
                text: loading ? qsTr("Unlocking") : qsTr("Unlock")

                Layout.fillWidth: true
                Layout.topMargin: 24

                onClicked: {
                    var err = false

                    if (secondPasswordTextField.text.length == 0) {
                        secondPasswordTextField.error = true
                        secondPasswordTextField.assistiveText = qsTr("Enter username or email")
                        err = true
                    } else {
                        secondPasswordTextField.error = false
                        secondPasswordTextField.assistiveText = qsTr("")
                    }

                    if (err) {
                        return
                    }

                    secondPasswordTextField.enabled = false

                    enabled = false
                    loading = true

                    if (root.user !== undefined) {
                        root.user.login2Password(usernameTextField.text, secondPasswordTextField.text)
                        return
                    }

                    root.login2Password(usernameTextField.text, secondPasswordTextField.text)
                }
            }
        }
    }

    states: [
        State {
            name: "Page 1"
            PropertyChanges {
                target: stackLayout
                currentIndex: 0
            }
        },
        State {
            name: "Page 2"
            PropertyChanges {
                target: stackLayout
                currentIndex: 1
            }
        },
        State {
            name: "Page 3"
            PropertyChanges {
                target: stackLayout
                currentIndex: 2
            }
        }
    ]
}