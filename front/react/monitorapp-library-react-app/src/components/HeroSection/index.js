import React, { useState } from "react";
import Video from "../../videos/onlywaves.mp4";
import { Button } from "../ButtonElements";
import {
    HeroContainer,
    HeroBg,
    VideoBg,
    HeroContent,
    HeroH1,
    HeroP,
    HeroBtnWrapper,
    ArrowFoward,
    ArrowRight,
} from "./HeroElements";

const HeroSection = () => {
    const [hover, setHover] = useState(false);
    const onHover = () => {
        setHover(!hover);
    };

    return (
        <HeroContainer>
            <HeroBg>
                <VideoBg autoPlay loop muted src={Video} type="vedio/mp4" />
            </HeroBg>
            <HeroContent>
                <HeroH1>MonitorApp Library</HeroH1>
                <HeroP>
                    Find, Discover and Improve yourself with monitorapp library.
                    we provide various genre, even comics!
                    <br />
                    Just come and read. You can feel difference.
                </HeroP>
                <HeroBtnWrapper>
                    <Button
                        to="about"
                        onMouseEnter={onHover}
                        onMouseLeave={onHover}
                        primary="true"
                        dark="true"

                        smooth={true}
                        duration={500}
                        spy={true}
                        exact="true"
                        offset={-80}
                    >
                        Find Out {hover ? <ArrowFoward /> : <ArrowRight />}
                    </Button>
                </HeroBtnWrapper>
            </HeroContent>
        </HeroContainer>
    );
};

export default HeroSection;
