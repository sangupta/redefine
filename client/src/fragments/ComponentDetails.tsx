/**
 *
 * Redefine - UI component documentation
 * 
 * MIT License.
 * Copyright (c) 2022, Sandeep Gupta.
 * 
 * Use of this source code is governed by a MIT style license
 * that can be found in LICENSE file in the code repository.
 * 
 **/

import StyledMarkdown from "./StyledMarkdown";

export default function ComponentDetails({ component }: { component: ComponentDef }) {
    return <StyledMarkdown>{component.description}</StyledMarkdown>
}
