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

import styled from 'styled-components';
import ReactMarkdown from 'react-markdown';

const StyledMarkdown = styled(ReactMarkdown)`
    line-height: 24px;

    & code {
        color: rgb(214, 51, 132);
        background-color: var(--redefine-bg);
    }

    & p {
        margin: 0;
    }
`;

export default StyledMarkdown;
